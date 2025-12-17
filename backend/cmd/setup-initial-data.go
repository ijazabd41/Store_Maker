package main

import (
	"log"

	"storemaker-backend/config"
	"storemaker-backend/database"
	"storemaker-backend/models"
	"storemaker-backend/utils"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.Initialize(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Creating admin user...")
	
	// Check if admin already exists
	var existingAdmin models.User
	if err := db.Where("role = ?", models.RoleAdmin).First(&existingAdmin).Error; err == nil {
		log.Printf("Admin user already exists: %s\n", existingAdmin.Email)
	} else {
		// Create admin user
		hashedPassword, err := utils.HashPassword("admin123")
		if err != nil {
			log.Fatal("Failed to hash password:", err)
		}

		admin := models.User{
			Email:     "admin@storemaker.com",
			Password:  hashedPassword,
			FirstName: "Admin",
			LastName:  "User",
			Role:      models.RoleAdmin,
			IsActive:  true,
		}

		if err := db.Create(&admin).Error; err != nil {
			log.Fatal("Failed to create admin user:", err)
		}
		log.Println("✓ Admin user created successfully")
		log.Println("  Email: admin@storemaker.com")
		log.Println("  Password: admin123")
	}

	log.Println("\nCreating templates...")
	
	// Check if templates already exist
	var count int64
	db.Model(&models.Template{}).Count(&count)
	if count > 0 {
		log.Printf("Templates already exist (count: %d)\n", count)
	} else {
		templates := []models.Template{
			{
				Name:        "Fashion Store",
				Description: "Modern fashion and apparel store template with clean design",
				Category:    models.TemplateCategoryFashion,
				PreviewURL:  "/templates/fashion-preview.jpg",
				ThumbnailURL: "/templates/fashion-thumb.jpg",
				Config:      models.TemplateConfig{"colors": map[string]interface{}{"primary": "#000000", "secondary": "#ffffff", "accent": "#e5e7eb"}, "layout": "modern"},
				IsActive:    true,
				IsPremium:   false,
			},
			{
				Name:        "Electronics Hub",
				Description: "Tech-focused template perfect for electronics stores",
				Category:    models.TemplateCategoryElectronics,
				PreviewURL:  "/templates/electronics-preview.jpg",
				ThumbnailURL: "/templates/electronics-thumb.jpg",
				Config:      models.TemplateConfig{"colors": map[string]interface{}{"primary": "#2563eb", "secondary": "#1e40af", "accent": "#3b82f6"}, "layout": "tech"},
				IsActive:    true,
				IsPremium:   false,
			},
			{
				Name:        "Food & Beverage",
				Description: "Delicious template for restaurants and food businesses",
				Category:    models.TemplateCategoryFood,
				PreviewURL:  "/templates/food-preview.jpg",
				ThumbnailURL: "/templates/food-thumb.jpg",
				Config:      models.TemplateConfig{"colors": map[string]interface{}{"primary": "#dc2626", "secondary": "#991b1b", "accent": "#ef4444"}, "layout": "cozy"},
				IsActive:    true,
				IsPremium:   false,
			},
			{
				Name:        "Beauty & Cosmetics",
				Description: "Elegant template for beauty and cosmetics stores",
				Category:    models.TemplateCategoryBeauty,
				PreviewURL:  "/templates/beauty-preview.jpg",
				ThumbnailURL: "/templates/beauty-thumb.jpg",
				Config:      models.TemplateConfig{"colors": map[string]interface{}{"primary": "#ec4899", "secondary": "#db2777", "accent": "#f472b6"}, "layout": "elegant"},
				IsActive:    true,
				IsPremium:   false,
			},
			{
				Name:        "Home & Garden",
				Description: "Comfortable template for home and garden products",
				Category:    models.TemplateCategoryHome,
				PreviewURL:  "/templates/home-preview.jpg",
				ThumbnailURL: "/templates/home-thumb.jpg",
				Config:      models.TemplateConfig{"colors": map[string]interface{}{"primary": "#10b981", "secondary": "#059669", "accent": "#34d399"}, "layout": "natural"},
				IsActive:    true,
				IsPremium:   false,
			},
			{
				Name:        "Minimal Store",
				Description: "Clean and minimal design for any type of store",
				Category:    models.TemplateCategoryGeneral,
				PreviewURL:  "/templates/minimal-preview.jpg",
				ThumbnailURL: "/templates/minimal-thumb.jpg",
				Config:      models.TemplateConfig{"colors": map[string]interface{}{"primary": "#6366f1", "secondary": "#4f46e5", "accent": "#818cf8"}, "layout": "minimal"},
				IsActive:    true,
				IsPremium:   false,
			},
		}

		for _, template := range templates {
			if err := db.Create(&template).Error; err != nil {
				log.Printf("Failed to create template %s: %v\n", template.Name, err)
			} else {
				log.Printf("✓ Created template: %s\n", template.Name)
			}
		}
	}

	log.Println("\n✓ Initial data setup completed successfully!")
}

