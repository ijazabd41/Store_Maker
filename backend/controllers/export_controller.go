package controllers

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"storemaker-backend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExportController struct {
	db *gorm.DB
}

func NewExportController(db *gorm.DB) *ExportController {
	return &ExportController{db: db}
}

// DownloadStoreSource generates and downloads the complete store source code as a zip file
func (ctrl *ExportController) DownloadStoreSource(c *gin.Context) {
	storeIDStr := c.Param("id")
	storeID, err := strconv.ParseUint(storeIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Verify store ownership
	var store models.Store
	if err := ctrl.db.Where("id = ? AND owner_id = ?", storeID, userID).
		Preload("Theme").
		Preload("Pages").
		Preload("Products").
		First(&store).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found or you don't have permission"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch store"})
		}
		return
	}

	// Load store layout
	var storeLayout []models.Section
	ctrl.db.Where("store_id = ? AND page_id IS NULL", storeID).
		Preload("Components").
		Find(&storeLayout)

	// Load page layouts
	pageLayouts := make(map[uint][]models.Section)
	for _, page := range store.Pages {
		var sections []models.Section
		ctrl.db.Where("page_id = ?", page.ID).
			Preload("Components").
			Find(&sections)
		pageLayouts[page.ID] = sections
	}

	// Create a buffer to write the zip file
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// Generate and add files to zip
	if err := ctrl.generateStoreFiles(zipWriter, &store, storeLayout, pageLayouts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate store files"})
		return
	}

	// Close the zip writer
	if err := zipWriter.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create zip file"})
		return
	}

	// Set headers for file download
	filename := fmt.Sprintf("%s-store-source-%s.zip", store.Slug, time.Now().Format("2006-01-02"))
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Length", strconv.Itoa(buf.Len()))

	// Send the zip file
	c.Data(http.StatusOK, "application/zip", buf.Bytes())
}

func (ctrl *ExportController) generateStoreFiles(zipWriter *zip.Writer, store *models.Store, storeLayout []models.Section, pageLayouts map[uint][]models.Section) error {
	// 1. Create README.md
	readmeContent := ctrl.generateReadme(store)
	if err := ctrl.addFileToZip(zipWriter, "README.md", []byte(readmeContent)); err != nil {
		return err
	}

	// 2. Create store-config.json
	configContent, err := ctrl.generateStoreConfig(store)
	if err != nil {
		return err
	}
	if err := ctrl.addFileToZip(zipWriter, "store-config.json", configContent); err != nil {
		return err
	}

	// 3. Create theme.json
	if store.Theme != nil {
		themeContent, err := json.MarshalIndent(store.Theme, "", "  ")
		if err != nil {
			return err
		}
		if err := ctrl.addFileToZip(zipWriter, "theme.json", themeContent); err != nil {
			return err
		}
	}

	// 4. Create products.json
	productsContent, err := json.MarshalIndent(store.Products, "", "  ")
	if err != nil {
		return err
	}
	if err := ctrl.addFileToZip(zipWriter, "data/products.json", productsContent); err != nil {
		return err
	}

	// 5. Create pages data
	for _, page := range store.Pages {
		pageData := map[string]interface{}{
			"id":              page.ID,
			"title":           page.Title,
			"slug":            page.Slug,
			"type":            page.Type,
			"content":         page.Content,
			"seo_title":       page.SeoTitle,
			"seo_description": page.SeoDesc,
			"is_published":    page.IsPublished,
		}

		// Add page layout if exists
		if sections, ok := pageLayouts[page.ID]; ok {
			components := []map[string]interface{}{}
			for _, section := range sections {
				for _, component := range section.Components {
					components = append(components, map[string]interface{}{
						"type":  component.Type,
						"props": component.Config,
						"order": component.Order,
					})
				}
			}
			pageData["layout"] = components
		}

		pageContent, err := json.MarshalIndent(pageData, "", "  ")
		if err != nil {
			return err
		}
		filename := fmt.Sprintf("pages/%s.json", page.Slug)
		if err := ctrl.addFileToZip(zipWriter, filename, pageContent); err != nil {
			return err
		}
	}

	// 6. Create store layout
	storeLayoutData := []map[string]interface{}{}
	for _, section := range storeLayout {
		for _, component := range section.Components {
			storeLayoutData = append(storeLayoutData, map[string]interface{}{
				"type":  component.Type,
				"props": component.Config,
				"order": component.Order,
			})
		}
	}
	layoutContent, err := json.MarshalIndent(storeLayoutData, "", "  ")
	if err != nil {
		return err
	}
	if err := ctrl.addFileToZip(zipWriter, "store-layout.json", layoutContent); err != nil {
		return err
	}

	// 7. Create package.json for deployment
	packageJSON := ctrl.generatePackageJSON(store)
	if err := ctrl.addFileToZip(zipWriter, "package.json", []byte(packageJSON)); err != nil {
		return err
	}

	// 8. Create deployment guide
	deployGuide := ctrl.generateDeploymentGuide(store)
	if err := ctrl.addFileToZip(zipWriter, "DEPLOYMENT.md", []byte(deployGuide)); err != nil {
		return err
	}

	return nil
}

func (ctrl *ExportController) addFileToZip(zipWriter *zip.Writer, filename string, content []byte) error {
	fileWriter, err := zipWriter.Create(filename)
	if err != nil {
		return err
	}
	_, err = fileWriter.Write(content)
	return err
}

func (ctrl *ExportController) generateReadme(store *models.Store) string {
	return fmt.Sprintf(`# %s - Store Source Code

## Description
%s

## Store Information
- **Name:** %s
- **Slug:** %s
- **Status:** %s
- **URL:** %s.storemaker.com

## Contents
This archive contains:
- **store-config.json**: Complete store configuration
- **theme.json**: Theme and styling configuration
- **data/products.json**: All products data
- **pages/**: Individual page configurations and layouts
- **store-layout.json**: Main store layout configuration
- **package.json**: Dependencies for deployment
- **DEPLOYMENT.md**: Deployment instructions

## Getting Started

### Option 1: Re-import to StoreMaker
You can re-import this configuration back into StoreMaker platform.

### Option 2: Deploy as Static Site
1. Extract this archive
2. Run: npm install
3. Configure your deployment platform
4. Deploy using the instructions in DEPLOYMENT.md

### Option 3: Use Data for Custom Implementation
Use the JSON files as a data source for your own custom implementation.

## File Structure
`, store.Name, store.Description, store.Name, store.Slug, store.Status, store.Slug) + `
- README.md - This file
- store-config.json - Store metadata and settings
- theme.json - Theme configuration
- store-layout.json - Main layout components
- data/
  - products.json - Product catalog
- pages/
  - home.json - Home page configuration
  - about.json - About page configuration
  - (other pages...)
- package.json - Node.js dependencies
- DEPLOYMENT.md - Deployment guide

## Support
For questions or support, visit https://storemaker.com/support

Generated: ` + time.Now().Format("2006-01-02 15:04:05")
}

func (ctrl *ExportController) generateStoreConfig(store *models.Store) ([]byte, error) {
	config := map[string]interface{}{
		"id":          store.ID,
		"name":        store.Name,
		"slug":        store.Slug,
		"description": store.Description,
		"logo":        store.Logo,
		"favicon":     store.Favicon,
		"domain":      store.Domain,
		"subdomain":   store.Subdomain,
		"status":      store.Status,
		"template_id": store.TemplateID,
		"created_at":  store.CreatedAt,
		"updated_at":  store.UpdatedAt,
		"export_date": time.Now(),
		"version":     "1.0",
	}

	return json.MarshalIndent(config, "", "  ")
}

func (ctrl *ExportController) generatePackageJSON(store *models.Store) string {
	return fmt.Sprintf(`{
  "name": "%s-store",
  "version": "1.0.0",
  "description": "Exported store from StoreMaker",
  "main": "index.js",
  "scripts": {
    "build": "echo 'Build your static site here'",
    "deploy": "echo 'Deploy your site here'"
  },
  "keywords": ["store", "ecommerce", "storemaker"],
  "author": "StoreMaker",
  "license": "MIT",
  "dependencies": {
    "react": "^18.2.0",
    "react-dom": "^18.2.0"
  },
  "devDependencies": {
    "next": "^15.0.0"
  }
}`, store.Slug)
}

func (ctrl *ExportController) generateDeploymentGuide(store *models.Store) string {
	return "# Deployment Guide for " + store.Name + `

## Overview
This guide will help you deploy your store to various platforms.

## Prerequisites
- Node.js 18+ installed
- npm or yarn package manager
- Git (optional, but recommended)

## Deployment Options

### Option 1: Vercel (Recommended)
1. Install Vercel CLI: npm i -g vercel
2. Run: vercel
3. Follow the prompts
4. Your store will be live!

### Option 2: Netlify
1. Install Netlify CLI: npm i -g netlify-cli
2. Run: netlify init
3. Run: netlify deploy
4. Follow the prompts

### Option 3: GitHub Pages
1. Create a new GitHub repository
2. Push your code: 
   git init
   git add .
   git commit -m "Initial commit"
   git remote add origin <your-repo-url>
   git push -u origin main
3. Enable GitHub Pages in repository settings

### Option 4: Custom Server
1. Build your site: npm run build
2. Upload the out or build directory to your server
3. Configure your web server (Apache/Nginx) to serve the files

## Configuration Files

### Store Configuration
- **store-config.json**: Main store settings
- **theme.json**: Styling and theme
- **products.json**: Product catalog

### Customization
You can modify any JSON file to customize your store:
- Edit colors in theme.json
- Update products in products.json
- Modify page layouts in pages/*.json

## Environment Variables
Set these environment variables for production:

NEXT_PUBLIC_STORE_NAME=` + store.Name + `
NEXT_PUBLIC_STORE_SLUG=` + store.Slug + `
NEXT_PUBLIC_API_URL=<your-api-url>

## Domain Setup
1. Add your custom domain in your hosting platform
2. Update DNS records:
   - A record: Point to hosting platform IP
   - CNAME record: Point www to your domain

## Support
For help with deployment:
- Visit: https://storemaker.com/docs/deployment
- Email: support@storemaker.com
- Discord: https://discord.gg/storemaker

---
Generated: ` + time.Now().Format("2006-01-02 15:04:05")
}

