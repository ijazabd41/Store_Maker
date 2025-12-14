# Store Source Code Download Feature

## Overview
Users can now download their complete store source code as a ZIP file directly from the dashboard. This feature allows them to:
- Export all store data and configurations
- Deploy their store independently
- Backup their store content
- Use the data for custom implementations

## Implementation Details

### Backend

#### New Controller: `export_controller.go`
- **Location**: `backend/controllers/export_controller.go`
- **Main Function**: `DownloadStoreSource(c *gin.Context)`
- **Route**: `GET /api/v1/manage/stores/:id/download`

**What It Does:**
1. Authenticates the user and verifies store ownership
2. Fetches all store data (products, pages, layouts, theme)
3. Generates a ZIP file containing:
   - `README.md` - Store documentation
   - `store-config.json` - Store metadata
   - `theme.json` - Theme configuration
   - `store-layout.json` - Main layout
   - `data/products.json` - Product catalog
   - `pages/*.json` - Individual page configurations
   - `package.json` - Node.js dependencies
   - `DEPLOYMENT.md` - Deployment instructions
4. Returns the ZIP file for download

**Files Included in ZIP:**

```
store-name-source-YYYY-MM-DD.zip
├── README.md                    # Store overview and instructions
├── store-config.json            # Store metadata
├── theme.json                   # Theme and styling
├── store-layout.json            # Main store layout components
├── package.json                 # NPM dependencies
├── DEPLOYMENT.md                # Deployment guide
├── data/
│   └── products.json            # All products
└── pages/
    ├── home.json                # Home page config + layout
    ├── about.json               # About page config + layout
    └── (other pages...)         # Additional pages
```

**Security:**
- Requires authentication (JWT token)
- Verifies store ownership
- Only owner can download their store

### Frontend

#### API Method: `downloadStoreSource`
- **Location**: `frontend/src/lib/api.ts`
- **Signature**: `downloadStoreSource(storeId: number)`
- **Returns**: Blob response (ZIP file)

#### UI Implementation
- **Location**: `frontend/src/app/dashboard/stores/[id]/page.tsx`
- **Feature**: "Download Source Code" button in Quick Actions section
- **Styling**: Green gradient button with download icon
- **UX Features**:
  - Loading state during download
  - Toast notifications for success/error
  - Disabled state while downloading
  - Auto-generates filename: `{store-slug}-store-source-{date}.zip`

### Download Button Features:
```typescript
- Icon: Download arrow icon
- Text: "Download Source Code"
- Loading: "Downloading..."
- Color: Green gradient (from-green-600 to-emerald-600)
- Position: Quick Actions section
- Tooltip: "Download complete store source code as ZIP"
```

## Usage Instructions

### For Store Owners:

1. **Navigate to Store Management:**
   - Go to Dashboard → Your Store

2. **Click Download Button:**
   - Find "Download Source Code" in the Quick Actions section
   - Click the button

3. **Wait for Download:**
   - A loading toast will appear
   - The ZIP file will automatically download

4. **Extract and Use:**
   - Extract the ZIP file
   - Read `README.md` for instructions
   - See `DEPLOYMENT.md` for deployment options

### Deployment Options (Included in ZIP):

1. **Re-import to StoreMaker**
   - Use the JSON files to recreate the store

2. **Deploy as Static Site**
   - Vercel, Netlify, or GitHub Pages
   - Instructions included in DEPLOYMENT.md

3. **Custom Implementation**
   - Use JSON data as API responses
   - Build your own frontend

4. **Backup/Archive**
   - Keep as backup
   - Version control

## Technical Implementation

### Backend Route Registration:
```go
// backend/routes/routes.go
exportController := controllers.NewExportController(db)

storeRoutes.GET("/:id/download", exportController.DownloadStoreSource)
```

### Frontend Download Handler:
```typescript
const handleDownloadSource = async () => {
  const response = await api.downloadStoreSource(storeId)
  const blob = new Blob([response.data], { type: 'application/zip' })
  const url = window.URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = `${store.slug}-store-source-${date}.zip`
  link.click()
  window.URL.revokeObjectURL(url)
}
```

## ZIP File Contents Details

### README.md
- Store name and description
- File structure overview
- Getting started instructions
- Support information

### store-config.json
```json
{
  "id": 6,
  "name": "HelloStore",
  "slug": "hellostore",
  "description": "Here is description",
  "logo": "...",
  "favicon": "...",
  "status": "active",
  "template_id": 1,
  "created_at": "...",
  "updated_at": "...",
  "export_date": "...",
  "version": "1.0"
}
```

### theme.json
```json
{
  "colors": {
    "primary": "#10b981",
    "secondary": "#f0fdf4",
    "accent": "#22c55e",
    "background": "#ffffff",
    "text": "#0f172a"
  },
  "fonts": {
    "heading": "Inter",
    "body": "Inter"
  }
}
```

### pages/home.json
```json
{
  "id": 5,
  "title": "Home",
  "slug": "home",
  "type": "home",
  "content": "...",
  "is_published": true,
  "layout": [
    {
      "type": "hero-banner",
      "props": { ... },
      "order": 0
    },
    ...
  ]
}
```

### data/products.json
```json
[
  {
    "id": 1,
    "name": "Product Name",
    "slug": "product-slug",
    "price": 99.99,
    "description": "...",
    "images": [...],
    ...
  }
]
```

### DEPLOYMENT.md
- Deployment options (Vercel, Netlify, GitHub Pages, Custom Server)
- Step-by-step instructions
- Environment variables
- Domain setup guide

### package.json
- NPM package configuration
- Dependencies (React, Next.js)
- Build scripts

## Testing

### Backend Test:
```bash
# Login and get token
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@storemaker.com","password":"admin123"}'

# Download store source
curl -X GET http://localhost:8080/api/v1/manage/stores/6/download \
  -H "Authorization: Bearer YOUR_TOKEN" \
  --output store-source.zip
```

### Frontend Test:
1. Login to dashboard
2. Navigate to store management page
3. Click "Download Source Code" button
4. Verify ZIP file downloads
5. Extract and inspect contents

## Benefits

### For Users:
- ✅ Complete data ownership
- ✅ Easy backup and archival
- ✅ Deploy anywhere
- ✅ Portable between platforms
- ✅ Version control friendly

### For Platform:
- ✅ Increased user confidence
- ✅ No vendor lock-in concerns
- ✅ Professional feature
- ✅ Competitive advantage

## Future Enhancements

### Possible Additions:
1. **HTML Export**: Generate static HTML files
2. **Import Feature**: Re-import exported stores
3. **Scheduled Backups**: Automatic periodic exports
4. **Version History**: Track multiple export versions
5. **Partial Export**: Export only specific sections
6. **Format Options**: JSON, YAML, XML exports
7. **Direct Deploy**: One-click deploy to hosting platforms

## Restart Required

**Important:** The backend must be restarted to register the new route:
```bash
cd backend
go run cmd/main.go
```

## Files Modified/Created

### Backend:
- ✅ `backend/controllers/export_controller.go` (NEW)
- ✅ `backend/routes/routes.go` (MODIFIED)

### Frontend:
- ✅ `frontend/src/lib/api.ts` (MODIFIED)
- ✅ `frontend/src/app/dashboard/stores/[id]/page.tsx` (MODIFIED)

## Support

For issues or questions:
- Check the generated `README.md` in the ZIP file
- See `DEPLOYMENT.md` for deployment help
- Contact: support@storemaker.com

---
**Feature Status:** ✅ Complete - Ready for Testing
**Created:** 2025-12-14

