"use client"

import Link from 'next/link'
import { Store } from '@/types'
import {
  EyeIcon,
  PencilIcon,
  TrashIcon,
  ChartBarIcon,
  BuildingStorefrontIcon,
} from '@heroicons/react/24/outline'

interface StoreCardProps {
  store: Store
  onDelete: () => void
}

export function StoreCard({ store, onDelete }: StoreCardProps) {

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'active':
        return 'bg-green-100 text-green-800 dark:bg-green-900/20 dark:text-green-400'
      case 'inactive':
        return 'bg-red-100 text-red-800 dark:bg-red-900/20 dark:text-red-400'
      case 'draft':
        return 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/20 dark:text-yellow-400'
      default:
        return 'bg-gray-100 text-gray-800 dark:bg-gray-900/20 dark:text-gray-400'
    }
  }

  const formatDate = (dateString: string) => {
    if (!dateString) return 'Unknown date'
    
    try {
      const date = new Date(dateString)
      if (isNaN(date.getTime())) {
        return 'Invalid date'
      }
      return date.toLocaleDateString()
    } catch (error) {
      console.error('Date parsing error:', error)
      return 'Invalid date'
    }
  }

  const handleCardClick = (e: React.MouseEvent) => {
    // Don't navigate if clicking on action buttons
    if ((e.target as HTMLElement).closest('button, a')) {
      return
    }
    
    if (store.id) {
      window.location.href = `/dashboard/stores/${store.id}`
    }
  }

  return (
    <div 
      onClick={handleCardClick}
      className="group relative overflow-hidden rounded-2xl transition-all duration-300 hover:scale-[1.02] cursor-pointer"
    >
      {/* Gradient Background */}
      <div className="absolute inset-0 bg-gradient-to-br from-primary-50 via-white to-secondary-50 dark:from-dark-900 dark:via-dark-800 dark:to-dark-900 opacity-50"></div>
      
      {/* Animated gradient overlay on hover */}
      <div className="absolute inset-0 bg-gradient-to-br from-primary-500/10 via-transparent to-secondary-500/10 opacity-0 group-hover:opacity-100 transition-opacity duration-500"></div>
      
      {/* Card border glow effect */}
      <div className="absolute inset-0 rounded-2xl border-2 border-transparent bg-gradient-to-br from-primary-200 via-secondary-200 to-primary-200 dark:from-primary-900 dark:via-secondary-900 dark:to-primary-900 opacity-0 group-hover:opacity-20 transition-opacity duration-300"></div>
      
      <div className="relative bg-white/80 dark:bg-dark-900/80 backdrop-blur-sm rounded-2xl p-6 border border-gray-200/50 dark:border-dark-700/50 shadow-soft hover:shadow-elegant transition-all duration-300">
        {/* Header */}
        <div className="flex items-start justify-between mb-4">
          <div className="flex items-center min-w-0 flex-1">
            {store.logo ? (
              <div className="relative">
                <img
                  src={store.logo}
                  alt={store.name}
                  className="w-14 h-14 rounded-xl object-cover flex-shrink-0 ring-2 ring-primary-100 dark:ring-primary-900/50"
                />
                <div className="absolute -inset-0.5 bg-gradient-to-br from-primary-400 to-secondary-400 rounded-xl blur opacity-30 group-hover:opacity-50 transition-opacity"></div>
              </div>
            ) : (
              <div className="relative w-14 h-14 bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl flex items-center justify-center flex-shrink-0 shadow-lg">
                <BuildingStorefrontIcon className="h-7 w-7 text-white" />
                <div className="absolute inset-0 bg-gradient-to-br from-white/20 to-transparent rounded-xl"></div>
              </div>
            )}
            <div className="ml-4 min-w-0 flex-1">
              <h3 className="text-lg font-bold text-gray-900 dark:text-white truncate mb-1 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors">
                {store.name}
              </h3>
              <p className="text-sm text-gray-600 dark:text-gray-400 truncate flex items-center">
                <svg className="w-3.5 h-3.5 mr-1.5 opacity-70" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9" />
                </svg>
                {store.slug ? `${store.slug}.storemaker.com` : store.domain || store.subdomain || 'No domain set'}
              </p>
            </div>
          </div>
          
          <span className={`px-3 py-1.5 text-xs font-bold rounded-full flex-shrink-0 ml-3 ${getStatusColor(store.status || 'draft')} shadow-sm`}>
            {store.status || 'draft'}
          </span>
        </div>

        {/* Description */}
        <p className="text-gray-700 dark:text-gray-300 text-sm mb-5 line-clamp-2 leading-relaxed">
          {store.description || 'No description provided'}
        </p>

        {/* Stats */}
        <div className="flex items-center justify-between mb-5 pb-5 border-b border-gray-200 dark:border-dark-700">
          <div className="flex items-center space-x-4">
            <div className="flex items-center text-sm text-gray-600 dark:text-gray-400 bg-gray-100 dark:bg-dark-800 px-3 py-1.5 rounded-lg">
              <ChartBarIcon className="h-4 w-4 mr-1.5 text-primary-500" />
              <span className="font-medium">0 orders</span>
            </div>
            <div className="flex items-center text-sm text-gray-600 dark:text-gray-400 bg-gray-100 dark:bg-dark-800 px-3 py-1.5 rounded-lg">
              <svg className="h-4 w-4 mr-1.5 text-success-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <span className="font-medium">$0</span>
            </div>
          </div>
        </div>

        {/* Actions */}
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-2">
            {store.status === 'active' && store.slug && (
              <Link
                href={`/stores/${store.slug}`}
                target="_blank"
                className="p-2.5 text-gray-600 hover:text-primary-600 dark:text-gray-400 dark:hover:text-primary-400 transition-all rounded-xl hover:bg-primary-50 dark:hover:bg-primary-900/20 hover:shadow-sm"
                title="View Live Store"
                onClick={(e) => e.stopPropagation()}
              >
                <EyeIcon className="h-5 w-5" />
              </Link>
            )}
            
            <Link
              href={store.id ? `/dashboard/stores/${store.id}` : '#'}
              className={`p-2.5 transition-all rounded-xl ${
                store.id 
                  ? 'text-gray-600 hover:text-primary-600 dark:text-gray-400 dark:hover:text-primary-400 hover:bg-primary-50 dark:hover:bg-primary-900/20 hover:shadow-sm'
                  : 'text-gray-300 cursor-not-allowed dark:text-gray-600'
              }`}
              title={store.id ? "Edit Store" : "Store data loading..."}
              onClick={(e) => {
                e.stopPropagation()
                if (!store.id) e.preventDefault()
              }}
            >
              <PencilIcon className="h-5 w-5" />
            </Link>
            
            <button
              onClick={(e) => {
                e.stopPropagation()
                if (store.id) onDelete()
              }}
              className={`p-2.5 transition-all rounded-xl ${
                store.id 
                  ? 'text-gray-600 hover:text-red-600 dark:text-gray-400 dark:hover:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 hover:shadow-sm'
                  : 'text-gray-300 cursor-not-allowed dark:text-gray-600'
              }`}
              title={store.id ? "Delete Store" : "Store data loading..."}
              disabled={!store.id}
            >
              <TrashIcon className="h-5 w-5" />
            </button>
          </div>

          <Link
            href={store.id ? `/dashboard/stores/${store.id}` : '#'}
            className={`group/btn flex items-center px-4 py-2.5 font-semibold rounded-xl transition-all duration-300 ${
              store.id 
                ? 'bg-gradient-to-r from-primary-600 to-secondary-600 hover:from-primary-700 hover:to-secondary-700 text-white shadow-md hover:shadow-lg hover:shadow-primary-500/50'
                : 'bg-gray-300 text-gray-500 cursor-not-allowed'
            }`}
            onClick={(e) => {
              e.stopPropagation()
              if (!store.id) e.preventDefault()
            }}
          >
            <span className="text-sm">Manage</span>
            <svg className="w-4 h-4 ml-2 transition-transform group-hover/btn:translate-x-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 7l5 5m0 0l-5 5m5-5H6" />
            </svg>
          </Link>
        </div>

        {/* Created date - bottom corner */}
        <div className="mt-4 pt-4 border-t border-gray-200 dark:border-dark-700 flex items-center text-xs text-gray-500 dark:text-gray-400">
          <svg className="w-4 h-4 mr-1.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          Created {formatDate(store.created_at)}
        </div>
      </div>
    </div>
  )
}