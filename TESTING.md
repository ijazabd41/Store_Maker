# StoreMaker Testing Guide

## Quick Test Summary

✅ **Backend Status**: Running on `http://localhost:8080`  
✅ **Frontend Status**: Running on `http://localhost:3000`  
✅ **Database**: Connected and migrated  
✅ **Admin User**: Created (admin@storemaker.com / admin123)  
✅ **Templates**: 6 templates seeded successfully

## API Endpoints Testing

### Health Check
```bash
curl http://localhost:8080/health
```
Expected: `{"status":"healthy","message":"Storemaker Backend API is running"}`

### Authentication

#### Register New User
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe"
  }'
```

#### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

Save the `access_token` from the response for authenticated requests.

### Templates

#### Get All Templates
```bash
curl http://localhost:8080/api/v1/templates
```

#### Get Single Template
```bash
curl http://localhost:8080/api/v1/templates/1
```

### Stores (Requires Authentication)

#### Get User Stores
```bash
curl http://localhost:8080/api/v1/manage/stores \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

#### Create Store
```bash
curl -X POST http://localhost:8080/api/v1/manage/stores \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Test Store",
    "description": "A wonderful store",
    "template_id": 1
  }'
```

#### Get Store by Slug (Public)
```bash
curl http://localhost:8080/api/v1/stores/my-test-store
```

### Products (Requires Authentication)

#### Create Product
```bash
curl -X POST http://localhost:8080/api/v1/manage/stores/1/products \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Product",
    "description": "Product description",
    "short_description": "Short description",
    "price": 29.99,
    "sku": "TEST-001",
    "stock": 100,
    "images": ["https://example.com/image.jpg"]
  }'
```

#### Get Products for Store
```bash
curl http://localhost:8080/api/v1/manage/stores/1/products \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

#### Get Public Store Products
```bash
curl http://localhost:8080/api/v1/stores/my-test-store/products
```

### File Upload

#### Upload Logo
```bash
curl -X POST http://localhost:8080/api/v1/manage/stores/1/logo \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -F "file=@/path/to/logo.png"
```

#### Upload Favicon
```bash
curl -X POST http://localhost:8080/api/v1/manage/stores/1/favicon \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -F "file=@/path/to/favicon.ico"
```

## Frontend Testing

### Manual Testing Checklist

#### 1. Landing Page (/)
- [ ] Page loads without errors
- [ ] Navigation bar displays correctly
- [ ] Mobile menu works
- [ ] All links are functional
- [ ] Sign In/Get Started buttons work
- [ ] Hero section displays properly
- [ ] Features section visible
- [ ] CTA buttons functional
- [ ] Footer displays correctly

#### 2. Authentication

**Register Page (/auth/register)**
- [ ] Form displays correctly
- [ ] Email validation works
- [ ] Password validation works (min 6 chars)
- [ ] Password visibility toggle works
- [ ] All fields required
- [ ] Shows error for existing email
- [ ] Success redirects to dashboard
- [ ] Login link works

**Login Page (/auth/login)**
- [ ] Form displays correctly
- [ ] Email validation works
- [ ] Password validation works
- [ ] Password visibility toggle works
- [ ] Remember me checkbox displays
- [ ] Forgot password link displays
- [ ] Invalid credentials show error
- [ ] Success redirects based on role
- [ ] Register link works

#### 3. Dashboard (/dashboard)
- [ ] Requires authentication
- [ ] Redirects to login if not authenticated
- [ ] Shows user stores
- [ ] Stats display correctly (Total, Active, Draft)
- [ ] Create Store button works
- [ ] Store cards display properly
- [ ] Empty state shows when no stores
- [ ] Mobile responsive

**Create Store Modal**
- [ ] Opens when clicking Create Store
- [ ] Template selection works
- [ ] Form validation works
- [ ] Name field required
- [ ] Template preview displays
- [ ] AI mode toggle works (if AI configured)
- [ ] Create button submits correctly
- [ ] Success shows new store
- [ ] Close button works
- [ ] Cancel button works

#### 4. Store Management (/dashboard/stores/[id])
- [ ] Store details display
- [ ] Edit store information
- [ ] Upload logo works
- [ ] Upload favicon works
- [ ] Navigation to products works
- [ ] Navigation to pages works
- [ ] Navigation to builder works
- [ ] Settings accessible

#### 5. Store Builder (/dashboard/stores/[id]/builder)
- [ ] Builder interface loads
- [ ] Component library displays
- [ ] Drag and drop works
- [ ] Preview updates in real-time
- [ ] Save layout works
- [ ] Theme customizer works
- [ ] Undo/Redo works (if implemented)
- [ ] Mobile preview works

#### 6. Products Management (/dashboard/stores/[id]/products)
- [ ] Products list displays
- [ ] Add product button works
- [ ] Product form validation
- [ ] Image upload works
- [ ] Edit product works
- [ ] Delete product works (with confirmation)
- [ ] Stock management works
- [ ] SKU generation/validation

#### 7. Pages Management (/dashboard/stores/[id]/pages)
- [ ] Pages list displays
- [ ] Create page button works
- [ ] Page types work (home, about, contact, etc.)
- [ ] Edit page works
- [ ] Delete page works
- [ ] Publish/Unpublish works
- [ ] Page builder integration

#### 8. Public Store View (/stores/[slug])
- [ ] Store loads with correct branding
- [ ] Logo/Favicon display
- [ ] Navigation works
- [ ] Product catalog displays
- [ ] Product filtering/search works (if implemented)
- [ ] Cart functionality
- [ ] Checkout process
- [ ] Newsletter subscription works
- [ ] Responsive on mobile
- [ ] Theme applies correctly

#### 9. Admin Panel (/admin) - Admin users only
- [ ] Requires admin role
- [ ] All stores list
- [ ] Store status management
- [ ] Template management
- [ ] Analytics dashboard
- [ ] User management (if implemented)

## Automated Testing

### Backend Tests

```bash
cd backend
go test ./... -v
```

### Frontend Tests

```bash
cd frontend
npm run test
npm run lint
```

## Performance Testing

### Load Testing (using Apache Bench)

```bash
# Test health endpoint
ab -n 1000 -c 10 http://localhost:8080/health

# Test templates endpoint
ab -n 500 -c 5 http://localhost:8080/api/v1/templates

# Test authenticated endpoint
ab -n 100 -c 2 -H "Authorization: Bearer YOUR_TOKEN" http://localhost:8080/api/v1/manage/stores
```

### Frontend Performance

1. Open Chrome DevTools
2. Run Lighthouse audit
3. Check for:
   - Performance score
   - Accessibility score
   - Best practices
   - SEO score

Expected scores:
- Performance: 80+
- Accessibility: 90+
- Best Practices: 90+
- SEO: 90+

## Browser Compatibility

Test on:
- [ ] Chrome (latest)
- [ ] Firefox (latest)
- [ ] Safari (latest)
- [ ] Edge (latest)
- [ ] Mobile Safari (iOS)
- [ ] Mobile Chrome (Android)

## Security Testing

### XSS Prevention
- [ ] User inputs are sanitized
- [ ] HTML entities escaped
- [ ] No eval() or innerHTML with user data

### CSRF Protection
- [ ] API uses proper authentication headers
- [ ] No sensitive operations via GET requests

### SQL Injection
- [ ] All queries use parameterized statements (GORM handles this)
- [ ] No raw SQL with user input

### Authentication
- [ ] JWT tokens expire correctly
- [ ] Refresh tokens work
- [ ] Logout clears tokens
- [ ] Protected routes require auth
- [ ] Role-based access control works

## Known Issues and Limitations

### Current Limitations:
1. AI store creation requires OpenAI API key (optional feature)
2. Email functionality not yet implemented
3. Payment processing not yet integrated
4. Real-time features not yet implemented

### Planned Improvements:
1. Add comprehensive test suite
2. Implement email notifications
3. Add payment gateway integration
4. Real-time order updates (WebSockets)
5. Advanced analytics dashboard
6. Multi-language support
7. SEO optimization tools

## Bug Reporting

When reporting bugs, please include:
1. Steps to reproduce
2. Expected behavior
3. Actual behavior
4. Browser/device information
5. Screenshots if applicable
6. Console errors (if any)

## Test Coverage Goals

- Backend: 80%+ coverage
- Frontend: 70%+ coverage
- Critical paths: 100% coverage

## Continuous Integration

### GitHub Actions Workflow Example:

```yaml
name: CI

on: [push, pull_request]

jobs:
  backend:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: 1.21
      - run: cd backend && go test ./...
      
  frontend:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-node@v2
        with:
          node-version: 18
      - run: cd frontend && npm install && npm run build
```

## Post-Deployment Verification

After deploying to production, verify:
- [ ] Health endpoint responds
- [ ] HTTPS works correctly
- [ ] Database connection established
- [ ] File uploads work
- [ ] Email sending works (when implemented)
- [ ] All critical user flows work
- [ ] Performance is acceptable
- [ ] Error tracking configured
- [ ] Backups are running
- [ ] Monitoring is active

## Support and Maintenance

Regular maintenance tasks:
1. Monitor error logs daily
2. Review performance metrics weekly
3. Update dependencies monthly
4. Backup verification weekly
5. Security updates as needed
6. User feedback review weekly

