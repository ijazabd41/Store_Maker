# 🛍️ StoreMaker User Manual

## Welcome to StoreMaker

StoreMaker is a comprehensive e-commerce platform that empowers entrepreneurs, small businesses, and agencies to create stunning online stores without coding. Built with modern technologies and featuring AI-powered store creation, StoreMaker offers everything you need to launch and manage successful online businesses.

---

## 📋 Table of Contents

1. [Getting Started](#getting-started)
2. [User Roles & Permissions](#user-roles--permissions)
3. [Creating Your First Store](#creating-your-first-store)
4. [Store Builder & Customization](#store-builder--customization)
5. [Product Management](#product-management)
6. [Order Management](#order-management)
7. [Store Settings](#store-settings)
8. [Admin Panel](#admin-panel)
9. [Technical Specifications](#technical-specifications)
10. [API Reference](#api-reference)
11. [Troubleshooting](#troubleshooting)

---

## 🚀 Getting Started

### Prerequisites

Before you begin, ensure you have:
- **Web Browser**: Modern browser (Chrome, Firefox, Safari, Edge)
- **Internet Connection**: Stable broadband connection
- **Device**: Desktop, tablet, or mobile device

### Account Registration

1. Visit the StoreMaker homepage
2. Click **"Get Started"** or **"Sign Up"**
3. Fill in your registration details:
   - Email address
   - Password (minimum 8 characters)
   - First name and last name
4. Verify your email address
5. Log in to your account

### Default Admin Credentials

For testing purposes, you can use:
- **Email**: `admin@storemaker.com`
- **Password**: `admin123`

---

## 👥 User Roles & Permissions

### Customer
- Browse and purchase products
- Create account for order tracking
- Manage personal profile

### Merchant (Store Owner)
- Create and manage multiple stores
- Full control over store content and settings
- Product and order management
- Access to analytics and reports

### Admin (System Administrator)
- Manage all stores in the system
- User management
- System-wide analytics
- Template management

---

## 🏪 Creating Your First Store

### Method 1: Template-Based Creation

1. Log in to your StoreMaker dashboard
2. Click **"Create Store"**
3. Choose from 6 pre-built templates:
   - **Fashion**: Clothing and accessories
   - **Electronics**: Tech products and gadgets
   - **Food & Beverage**: Restaurant and food delivery
   - **Beauty**: Cosmetics and skincare
   - **Home & Garden**: Furniture and decor
   - **Minimal**: Clean, professional design

4. Customize basic information:
   - Store name
   - Description
   - Logo and favicon (optional)
5. Click **"Create Store"**

### Method 2: AI-Powered Store Creation

1. Select **"AI Store Creation"**
2. Provide business details:
   - Business type/industry
   - Target audience
   - Color preferences
   - Brand style
3. AI analyzes your requirements and generates:
   - Optimized store layout
   - Suggested color schemes
   - Content recommendations
   - Component placement

### Store Status Management

- **Draft**: Private, under development
- **Active**: Public, live store
- **Inactive**: Temporarily disabled

---

## 🎨 Store Builder & Customization

### Visual Store Builder

StoreMaker features a powerful drag-and-drop builder with 20+ pre-built components:

#### Hero Sections
- **Banner**: Large hero with text and image
- **Split**: Two-column hero layout
- **Video**: Video background hero
- **Minimal**: Clean text-only hero

#### Content Components
- **Text Blocks**: Rich text content
- **Image Galleries**: Photo showcases
- **Video Players**: Embedded videos
- **Testimonials**: Customer reviews
- **FAQs**: Frequently asked questions

#### Interactive Elements
- **Product Carousels**: Rotating product displays
- **Feature Grids**: Icon-based feature lists
- **Contact Forms**: Customer inquiry forms
- **Social Proof**: Trust indicators

### Page Management

#### Page Types
- **Home Page**: Main landing page
- **About Page**: Business story and information
- **Contact Page**: Contact form and details
- **Custom Pages**: Unlimited additional pages

#### Page Builder Features
- Section-based organization
- Drag-and-drop component placement
- Real-time preview
- SEO optimization fields

### Theme Customization

#### Color Palette
- Primary, secondary, and accent colors
- Text and background colors
- Custom color picker

#### Typography
- Heading and body font selection
- Font size and weight controls
- Custom font uploads (planned)

#### Layout Options
- Header styles (sticky, transparent, centered)
- Footer designs
- Button styles (rounded, square, outlined)

### Real-time Preview

- Desktop and mobile preview modes
- Live changes as you build
- Responsive design testing

---

## 🛍️ Product Management

### Product Catalog

#### Adding Products

1. Navigate to your store dashboard
2. Go to **"Products"** section
3. Click **"Add Product"**
4. Fill in product details:
   - Name and description
   - Price and compare-at price
   - SKU and inventory
   - Product images (up to 10)
   - Categories and tags

#### Product Variants

- Size, color, and custom options
- Individual pricing per variant
- Separate inventory tracking
- SKU management per variant

#### Advanced Features

- **Digital Products**: Downloadable files
- **SEO Optimization**: Meta titles and descriptions
- **Rich Descriptions**: HTML content support
- **Product Categories**: Hierarchical organization

### Inventory Management

- Real-time stock tracking
- Low stock alerts (planned)
- Automatic out-of-stock handling
- Bulk inventory updates

### Product Display

- Grid and list view options
- Smart filtering and sorting
- Search functionality
- Quick view modals
- Related product suggestions

---

## 📦 Order Management

### Order Processing

#### Order Status Workflow
1. **Pending**: Initial order placement
2. **Confirmed**: Payment verified
3. **Processing**: Order preparation
4. **Shipped**: Order dispatched
5. **Delivered**: Customer received
6. **Cancelled**: Order cancelled
7. **Refunded**: Refund processed

### Order Dashboard

- View all orders in one place
- Filter by status, date, customer
- Search by order number or email
- Export order reports (planned)

### Customer Management

- Customer profiles and order history
- Guest checkout support
- Customer communication tools (planned)
- Loyalty programs (planned)

---

## ⚙️ Store Settings

### General Settings

- Store name and description
- Contact information
- Business hours
- Social media links

### Commerce Settings

- Currency selection
- Tax configuration
- Shipping rates
- Guest checkout toggle
- Order numbering format

### Domain Configuration

- Custom domain support
- Automatic subdomain generation
- SSL certificate management

### Branding

- Logo and favicon upload
- Brand color scheme
- Email templates (planned)

---

## 👑 Admin Panel

### System Management

#### Store Oversight
- View all stores in the system
- Activate/deactivate stores
- Monitor store performance
- Access store analytics

#### Template Management
- Create new store templates
- Edit existing templates
- Delete unused templates
- Template categorization

#### User Management
- View all registered users
- Manage user roles
- User activity monitoring
- Account suspension tools

### System Analytics

- Platform-wide revenue tracking
- User registration statistics
- Popular templates and features
- System performance metrics

---

## 💻 Technical Specifications

### System Requirements

#### Backend
- **Go**: Version 1.21 or higher
- **Database**: PostgreSQL 13+
- **Memory**: 512MB minimum, 1GB recommended
- **Storage**: 1GB minimum for database and uploads

#### Frontend
- **Node.js**: Version 18 or higher
- **Memory**: 256MB minimum
- **Browser**: Modern browser with JavaScript enabled

### Technology Stack

#### Backend (Go/Gin)
- **Framework**: Gin web framework
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT tokens
- **File Storage**: Local file system
- **API**: RESTful JSON API

#### Frontend (Next.js)
- **Framework**: Next.js 15 with App Router
- **Styling**: Tailwind CSS
- **State Management**: React Context
- **Forms**: React Hook Form with Zod validation
- **UI Components**: Custom component library

#### Database Schema
- **Users**: Account management
- **Stores**: Store information and settings
- **Products**: Product catalog and variants
- **Orders**: Order processing and tracking
- **Pages**: CMS page management
- **Templates**: Store template system

### Performance Features

- Server-side rendering (SSR)
- Image optimization
- Code splitting
- Database query optimization
- CDN-ready architecture

### Security Features

- JWT-based authentication
- Password hashing (bcrypt)
- SQL injection protection
- XSS prevention
- CORS configuration
- Input validation and sanitization

---

## 🔌 API Reference

### Authentication Endpoints

```bash
POST /api/v1/auth/register     # User registration
POST /api/v1/auth/login        # User login
POST /api/v1/auth/refresh      # Refresh access token
POST /api/v1/auth/logout       # User logout
```

### Store Management

```bash
GET    /api/v1/manage/stores           # Get user stores
POST   /api/v1/manage/stores           # Create store
GET    /api/v1/manage/stores/:id       # Get store details
PUT    /api/v1/manage/stores/:id       # Update store
DELETE /api/v1/manage/stores/:id       # Delete store
```

### Product Management

```bash
GET    /api/v1/manage/stores/:id/products      # Get store products
POST   /api/v1/manage/stores/:id/products      # Create product
PUT    /api/v1/manage/stores/:id/products/:id  # Update product
DELETE /api/v1/manage/stores/:id/products/:id  # Delete product
```

### Order Management

```bash
GET    /api/v1/manage/stores/:id/orders        # Get store orders
PUT    /api/v1/manage/stores/:id/orders/:id    # Update order status
```

### Public Store Endpoints

```bash
GET /stores/:slug                    # Get public store
GET /stores/:slug/products           # Get store products
POST /stores/:slug/orders            # Create order
POST /stores/:slug/newsletter/subscribe  # Newsletter signup
```

### Admin Endpoints

```bash
GET  /api/v1/admin/stores           # Get all stores
PUT  /api/v1/admin/stores/:id/status # Change store status
GET  /api/v1/admin/analytics        # System analytics
```

### Software Construction Concepts API

StoreMaker includes educational Software Construction Concepts (SQC) endpoints:

```bash
POST /api/v1/sqc/query     # Execute StoreQL queries
POST /api/v1/sqc/validate  # Validate query syntax
POST /api/v1/sqc/tokenize  # Tokenize queries
GET  /api/v1/sqc/grammar   # Get StoreQL grammar
GET  /api/v1/sqc/examples  # Run educational examples
GET  /api/v1/sqc/docs      # Get documentation
```

---

## 🔧 Troubleshooting

### Common Issues

#### Database Connection Problems
**Symptoms**: Application won't start, error messages about database connection
**Solutions**:
1. Verify PostgreSQL is running
2. Check database credentials in `.env`
3. Ensure database exists
4. Check network connectivity

#### Port Conflicts
**Symptoms**: "Port already in use" error
**Solutions**:
1. Change port in environment variables
2. Kill processes using the port
3. Use different ports for development

#### Authentication Issues
**Symptoms**: Can't log in, token errors
**Solutions**:
1. Clear browser localStorage
2. Check JWT_SECRET configuration
3. Verify user credentials
4. Check token expiration

#### Store Builder Problems
**Symptoms**: Builder not loading, components not saving
**Solutions**:
1. Clear browser cache
2. Check network connectivity
3. Verify user permissions
4. Try different browser

#### File Upload Issues
**Symptoms**: Images/logos not uploading
**Solutions**:
1. Check file size limits
2. Verify supported file formats
3. Check server storage permissions
4. Ensure uploads directory exists

### Performance Issues

#### Slow Loading
- Enable browser caching
- Optimize images
- Check server resources
- Review database queries

#### High Memory Usage
- Monitor Go application memory
- Check for memory leaks
- Optimize database queries
- Consider server upgrades

### Debug Mode

#### Backend Debug
```bash
cd backend
go run cmd/main.go -debug
```

#### Frontend Debug
```bash
cd frontend
DEBUG=* npm run dev
```

---

## 📈 Advanced Features

### Analytics & Reporting

#### Store Analytics
- Sales performance tracking
- Popular products identification
- Customer behavior insights
- Conversion rate monitoring

#### Revenue Tracking
- Daily, weekly, monthly reports
- Product performance analysis
- Customer segmentation
- Trend analysis

### Marketing Tools

#### Newsletter Management
- Subscriber list management
- Automated campaigns (planned)
- Email template customization
- GDPR compliance

#### SEO Optimization
- Meta tag management
- Sitemap generation
- Schema markup
- Mobile-first indexing ready

### Integration Capabilities

#### Third-Party Services
- Payment gateways (planned)
- Shipping providers (planned)
- Email marketing (planned)
- Analytics platforms (planned)

#### API Access
- RESTful API for custom integrations
- Webhook support (planned)
- Custom plugin development (planned)

---

## 🎯 Best Practices

### Store Design
- Use high-quality images
- Maintain consistent branding
- Optimize for mobile devices
- Keep navigation simple
- Use clear call-to-action buttons

### Product Management
- Write detailed product descriptions
- Use consistent naming conventions
- Maintain accurate inventory
- Set competitive pricing
- Use high-quality product photos

### Customer Experience
- Enable guest checkout
- Provide clear shipping information
- Set realistic delivery expectations
- Offer multiple payment options
- Provide excellent customer support

### Security
- Use strong passwords
- Enable two-factor authentication (planned)
- Keep software updated
- Monitor for suspicious activity
- Backup data regularly

---

## 📞 Support & Resources

### Getting Help
- **Documentation**: Comprehensive guides and tutorials
- **Community Forum**: Connect with other users (planned)
- **Video Tutorials**: Step-by-step video guides (planned)
- **Support Ticket**: Direct support for premium users

### Learning Resources
- **Getting Started Guide**: Basic setup and usage
- **Advanced Tutorials**: Complex features and integrations
- **API Documentation**: Technical integration guides
- **Best Practices**: Optimization and performance tips

### Community
- **GitHub Repository**: Source code and issue tracking
- **Blog**: Updates, tips, and industry insights
- **Newsletter**: Product updates and announcements

---

## 🚀 What's Next

### Upcoming Features
- **Payment Gateway Integration**: Stripe, PayPal support
- **Advanced Analytics**: Detailed reporting and insights
- **Mobile App**: iOS and Android applications
- **Multi-language Support**: International expansion
- **Advanced Marketing**: Email campaigns and automation
- **Inventory Management**: Advanced stock tracking
- **POS Integration**: Point of sale system integration

### Enterprise Features (Future)
- **White-label Solution**: Custom branding options
- **Multi-vendor Marketplace**: Multiple sellers support
- **Advanced Permissions**: Granular access control
- **API v2**: GraphQL implementation
- **Advanced Integrations**: ERP and CRM systems

---

## 📄 License & Terms

StoreMaker is built with modern, open-source technologies and is designed for both personal and commercial use. The platform is licensed under the MIT License, allowing flexibility for customization and extension.

For enterprise deployments and custom development services, please contact our sales team.

---

## 🎉 Conclusion

StoreMaker empowers you to create professional, high-converting online stores without the complexity of traditional e-commerce platforms. Whether you're a small business owner, entrepreneur, or agency, StoreMaker provides the tools and flexibility you need to succeed online.

**Ready to start building?** Create your first store today and join thousands of successful merchants who trust StoreMaker for their online business needs.

**Happy selling! 🚀**