# StoreMaker Production Deployment Guide

## Prerequisites

- **PostgreSQL** 13+ installed and running
- **Go** 1.21+ installed  
- **Node.js** 18+ installed
- **Domain names** configured (for production)
- **SSL certificates** (for HTTPS in production)

## Backend Deployment

### 1. Environment Setup

```bash
cd backend
cp config.env.production config.env
```

Edit `config.env` and configure:
- `DATABASE_URL`: Your production PostgreSQL connection string
- `JWT_SECRET`: Generate a strong random secret (use `openssl rand -base64 64`)
- `SERVER_URL`: Your backend domain (e.g., `https://api.yourdomain.com`)
- `FRONTEND_URL`: Your frontend domain for CORS

### 2. Database Setup

```bash
# Connect to PostgreSQL and create database
psql -U postgres
CREATE DATABASE storemaker;
\q

# Run initial data setup (creates admin user and templates)
go run cmd/setup-initial-data.go
```

**Default Admin Credentials:**
- Email: `admin@storemaker.com`
- Password: `admin123`
- **⚠️ IMPORTANT: Change this password immediately after first login!**

### 3. Build and Run

```bash
# Install dependencies
go mod download

# Build for production
go build -o bin/storemaker cmd/main.go

# Run the server
./bin/storemaker
```

Or use a process manager like `systemd`:

```ini
# /etc/systemd/system/storemaker.service
[Unit]
Description=StoreMaker Backend API
After=network.target postgresql.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/var/www/storemaker/backend
Environment="PATH=/usr/local/go/bin:/usr/bin"
ExecStart=/var/www/storemaker/backend/bin/storemaker
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl enable storemaker
sudo systemctl start storemaker
sudo systemctl status storemaker
```

## Frontend Deployment

### 1. Environment Setup

```bash
cd frontend
```

Create `.env.production` or `.env.local` with:
```
NEXT_PUBLIC_API_URL=https://api.yourdomain.com/api/v1
NEXT_PUBLIC_APP_URL=https://yourdomain.com
```

### 2. Build and Deploy

```bash
# Install dependencies
npm install

# Build for production
npm run build

# Start production server
npm start
```

Or deploy to Vercel/Netlify:

```bash
# Install Vercel CLI
npm i -g vercel

# Deploy
vercel --prod
```

## Nginx Configuration

### Backend API

```nginx
server {
    listen 80;
    server_name api.yourdomain.com;
    
    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
    
    location /uploads {
        alias /var/www/storemaker/backend/uploads;
        expires 30d;
        add_header Cache-Control "public, immutable";
    }
}
```

### Frontend

```nginx
server {
    listen 80;
    server_name yourdomain.com;
    
    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

## SSL Setup (Let's Encrypt)

```bash
# Install certbot
sudo apt install certbot python3-certbot-nginx

# Get SSL certificates
sudo certbot --nginx -d yourdomain.com -d www.yourdomain.com
sudo certbot --nginx -d api.yourdomain.com

# Auto-renewal
sudo certbot renew --dry-run
```

## Database Backup

```bash
# Create backup script
cat > /usr/local/bin/backup-storemaker.sh << 'EOF'
#!/bin/bash
BACKUP_DIR="/var/backups/storemaker"
DATE=$(date +%Y%m%d_%H%M%S)
mkdir -p $BACKUP_DIR

# Backup database
pg_dump -U postgres storemaker | gzip > $BACKUP_DIR/storemaker_$DATE.sql.gz

# Backup uploads
tar -czf $BACKUP_DIR/uploads_$DATE.tar.gz /var/www/storemaker/backend/uploads

# Keep only last 7 days
find $BACKUP_DIR -name "*.gz" -mtime +7 -delete
EOF

chmod +x /usr/local/bin/backup-storemaker.sh

# Add to crontab (daily at 2 AM)
(crontab -l 2>/dev/null; echo "0 2 * * * /usr/local/bin/backup-storemaker.sh") | crontab -
```

## Monitoring

### Health Check Endpoints

- Backend: `http://localhost:8080/health`
- Frontend: `http://localhost:3000`

### Logs

```bash
# Backend logs (if using systemd)
journalctl -u storemaker -f

# Nginx logs
tail -f /var/log/nginx/access.log
tail -f /var/log/nginx/error.log
```

## Security Checklist

- [ ] Changed default admin password
- [ ] Generated strong JWT_SECRET
- [ ] Configured SSL/HTTPS
- [ ] Set up firewall rules
- [ ] Configured database user with limited privileges
- [ ] Set up regular backups
- [ ] Configured rate limiting
- [ ] Set up monitoring and alerts
- [ ] Reviewed CORS settings
- [ ] Secured environment variables
- [ ] Set up log rotation

## Performance Optimization

### Database

```sql
-- Add indexes for better performance
CREATE INDEX idx_stores_owner_id ON stores(owner_id);
CREATE INDEX idx_stores_slug ON stores(slug);
CREATE INDEX idx_products_store_id ON products(store_id);
CREATE INDEX idx_orders_store_id ON orders(store_id);
CREATE INDEX idx_orders_customer_email ON orders(customer_email);
```

### Backend

- Enable gzip compression in Nginx
- Set up Redis for caching (optional)
- Use CDN for static assets

### Frontend

- Enable image optimization in Next.js
- Configure proper caching headers
- Use CDN for frontend assets (Vercel/CloudFlare)

## Troubleshooting

### Backend won't start

```bash
# Check logs
journalctl -u storemaker -n 50

# Check database connection
psql -U postgres -d storemaker -c "SELECT 1;"

# Check environment variables
env | grep DATABASE_URL
```

### Frontend build fails

```bash
# Clear cache
rm -rf .next node_modules
npm install
npm run build
```

### Database migration issues

```bash
# Check database migrations
go run cmd/main.go

# If needed, recreate database
dropdb storemaker
createdb storemaker
go run cmd/setup-initial-data.go
```

## Updating

### Backend

```bash
cd backend
git pull
go mod download
go build -o bin/storemaker cmd/main.go
sudo systemctl restart storemaker
```

### Frontend

```bash
cd frontend
git pull
npm install
npm run build
pm2 restart storemaker-frontend  # or your process manager
```

## Support

For issues and questions:
- Check logs first
- Review configuration
- Verify database connectivity
- Check firewall rules
- Ensure all services are running

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Next.js Deployment](https://nextjs.org/docs/deployment)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Nginx Documentation](https://nginx.org/en/docs/)

