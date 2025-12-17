# Bug Fix: Store Preview vs Published Page Mismatch

## Issue Description
The store page builder preview was showing different content than the published store page, even though both were supposed to display the same layout.

**Affected URLs:**
- Builder Preview: `http://localhost:3000/dashboard/stores/6/pages/5/builder`
- Published Store: `http://localhost:3000/stores/hellostore/pages/home`

## Root Cause
The home page component (`frontend/src/app/stores/[slug]/pages/home/page.tsx`) was loading the **store layout** instead of the **home page layout**.

### What Was Wrong:
```typescript
// INCORRECT - Loading store layout
const layoutResponse = await api.getPublicStoreLayout(slug)
```

### What It Should Be:
```typescript
// CORRECT - Loading home page layout
const layoutResponse = await api.getPublicPageLayout(slug, 'home')
```

## Technical Details

### Layout Architecture:
The application has two types of layouts:

1. **Store Layout** (`/stores/{slug}/layout`)
   - General store-wide layout settings
   - Used as a fallback when no page-specific layout exists

2. **Page Layout** (`/stores/{slug}/pages/{pageSlug}/layout`)
   - Specific to individual pages (home, about, contact, etc.)
   - Should be used for all custom page content

### The Problem:
- **Builder Preview** correctly loaded: `/api/v1/manage/stores/6/pages/5/layout` (Page Layout)
- **Public Home Page** incorrectly loaded: `/api/v1/stores/hellostore/layout` (Store Layout)
- This caused them to display different content

## Fix Applied

**File:** `frontend/src/app/stores/[slug]/pages/home/page.tsx`

**Changed Line 40 from:**
```typescript
const layoutResponse = await api.getPublicStoreLayout(slug)
```

**To:**
```typescript
const layoutResponse = await api.getPublicPageLayout(slug, 'home')
```

## Verification

### Backend API Test:
```bash
# Both endpoints now return the same data:

# Builder endpoint:
GET /api/v1/manage/stores/6/pages/5/layout
# Returns: {"components": [{"type": "hero-banner", "props": {"title": "Welcome to Our Store"}}...]}

# Public endpoint:
GET /api/v1/stores/hellostore/pages/home/layout
# Returns: {"components": [{"type": "hero-banner", "props": {"title": "Welcome to Our Store"}}...]}
```

## Impact
✅ The published store home page now displays the exact same content as the builder preview
✅ All page components (hero banners, product grids, etc.) render consistently
✅ Theme settings are applied correctly

## Related Files
- `frontend/src/app/stores/[slug]/pages/home/page.tsx` - Fixed
- `frontend/src/app/stores/[slug]/pages/[pageSlug]/page.tsx` - Already correct
- `frontend/src/lib/api.ts` - API definitions (no changes needed)
- `backend/controllers/customization_controller.go` - Backend endpoints (no changes needed)

## Testing Checklist
- [x] Verify builder preview loads page layout correctly
- [x] Verify published page loads page layout correctly
- [x] Verify page publishing status works
- [x] Verify both endpoints return identical data
- [x] Test theme customization applies to both views
- [x] Test component changes sync between builder and published page

## Additional Notes
- The page must be published (`is_published = true`) for it to appear in the public view
- The store layout endpoint (`getPublicStoreLayout`) is still valid for potential future use
- Custom pages (about, contact, etc.) were already using the correct endpoint

