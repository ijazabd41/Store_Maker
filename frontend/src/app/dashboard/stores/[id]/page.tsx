"use client"

import { useEffect, useState } from 'react'
import { useRouter, useParams } from 'next/navigation'
import { useAuth } from '@/contexts/auth-context'
import { api } from '@/lib/api'
import { Store } from '@/types'
import { StoreLayout } from '@/components/layout/store-layout'
import { ArrowLeftIcon } from '@heroicons/react/24/outline'
import toast from 'react-hot-toast'
import Link from 'next/link'

export default function StoreManagePage() {
  const { isLoading, isAuthenticated } = useAuth()
  const router = useRouter()
  const params = useParams()
  const storeId = params.id as string
  
  const [store, setStore] = useState<Store | null>(null)
  const [isLoadingStore, setIsLoadingStore] = useState(true)
  const [isDownloading, setIsDownloading] = useState(false)

  useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      router.push('/auth/login')
      return
    }

    if (isAuthenticated && storeId) {
      fetchStore()
    }
  }, [isLoading, isAuthenticated, router, storeId])

  const fetchStore = async () => {
    try {
      setIsLoadingStore(true)
      const response = await api.getStore(parseInt(storeId))
      setStore(response.data as Store)
    } catch (error) {
      console.error('Failed to fetch store:', error)
      toast.error('Failed to load store')
      router.push('/dashboard')
    } finally {
      setIsLoadingStore(false)
    }
  }

  const handleActivateStore = async () => {
    try {
      await api.updateStore(parseInt(storeId), { status: 'active' })
      toast.success('Store activated successfully!')
      fetchStore() // Refresh store data
    } catch (error) {
      console.error('Failed to activate store:', error)
      toast.error('Failed to activate store')
    }
  }

  const handleDeactivateStore = async () => {
    try {
      await api.updateStore(parseInt(storeId), { status: 'draft' })
      toast.success('Store deactivated successfully!')
      fetchStore() // Refresh store data
    } catch (error) {
      console.error('Failed to deactivate store:', error)
      toast.error('Failed to deactivate store')
    }
  }

  const handleDownloadSource = async () => {
    if (!store) return
    
    try {
      setIsDownloading(true)
      toast.loading('Preparing download...')
      
      const response = await api.downloadStoreSource(parseInt(storeId))
      
      // Create blob from response
      const blob = new Blob([response.data as BlobPart], { type: 'application/zip' })
      
      // Create download link
      const url = window.URL.createObjectURL(blob)
      const link = document.createElement('a')
      link.href = url
      link.download = `${store.slug}-store-source-${new Date().toISOString().split('T')[0]}.zip`
      
      // Trigger download
      document.body.appendChild(link)
      link.click()
      
      // Cleanup
      document.body.removeChild(link)
      window.URL.revokeObjectURL(url)
      
      toast.dismiss()
      toast.success('Store source code downloaded successfully!')
    } catch (error) {
      console.error('Failed to download store source:', error)
      toast.dismiss()
      toast.error('Failed to download store source code')
    } finally {
      setIsDownloading(false)
    }
  }

  if (isLoading || isLoadingStore) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="spinner h-8 w-8"></div>
      </div>
    )
  }

  if (!isAuthenticated || !store) {
    return null
  }

  return (
    <StoreLayout storeId={parseInt(storeId)} storeName={store.name}>
      <div className="space-y-6 sm:space-y-8">
        {/* Header */}
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between space-y-4 sm:space-y-0">
          <div className="space-y-2">
            <h1 className="text-2xl sm:text-3xl font-bold text-gray-900 dark:text-gray-100">
              Store Overview
            </h1>
            <p className="text-sm sm:text-base text-gray-600 dark:text-gray-400">
              Manage your store settings, products, and analytics
            </p>
          </div>
          <div className="flex flex-col sm:flex-row items-stretch sm:items-center space-y-3 sm:space-y-0 sm:space-x-4">
            <span className={`px-3 py-1 text-sm font-medium rounded-full text-center ${
              store.status === 'active' 
                ? 'bg-green-100 text-green-800 dark:bg-green-900/20 dark:text-green-400'
                : store.status === 'draft'
                ? 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/20 dark:text-yellow-400'
                : 'bg-red-100 text-red-800 dark:bg-red-900/20 dark:text-red-400'
            }`}>
              {store.status}
            </span>
            {/* Status Toggle Button */}
            {store.status === 'draft' ? (
              <button
                onClick={handleActivateStore}
                className="btn-primary w-full sm:w-auto"
              >
                Activate Store
              </button>
            ) : store.status === 'active' ? (
              <button
                onClick={handleDeactivateStore}
                className="btn-secondary w-full sm:w-auto"
              >
                Deactivate Store
              </button>
            ) : null}
            
            {store.status === 'active' && (
              <Link
                href={`/stores/${store.slug}`}
                target="_blank"
                className="btn-secondary w-full sm:w-auto text-center"
              >
                View Store
              </Link>
            )}
          </div>
        </div>

        {/* Draft Store Warning */}
        {store.status === 'draft' && (
          <div className="bg-yellow-50 dark:bg-yellow-900/20 border border-yellow-200 dark:border-yellow-800 rounded-lg p-4 mb-6">
            <div className="flex items-start">
              <svg className="w-5 h-5 text-yellow-600 dark:text-yellow-400 mr-3 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                <path fillRule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
              </svg>
              <div>
                <h3 className="text-sm font-medium text-yellow-800 dark:text-yellow-200">
                  Store is in Draft Mode
                </h3>
                <p className="text-sm text-yellow-700 dark:text-yellow-300 mt-1">
                  Your store is not publicly accessible. Click "Activate Store" to make it live.
                </p>
              </div>
            </div>
          </div>
        )}

        {/* Store Info */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6 lg:gap-8">
          <div className="lg:col-span-2">
            <div className="relative overflow-hidden rounded-2xl bg-white/80 dark:bg-dark-900/80 backdrop-blur-sm border border-gray-200/50 dark:border-dark-700/50 shadow-elegant">
              <div className="absolute top-0 left-0 right-0 h-32 bg-gradient-to-br from-primary-500 via-secondary-500 to-primary-600 opacity-10"></div>
              <div className="relative p-6">
                <div className="flex items-center mb-6">
                  <div className="p-2 bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl mr-3">
                    <svg className="h-6 w-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <h2 className="text-2xl font-bold bg-gradient-to-r from-primary-600 to-secondary-600 bg-clip-text text-transparent">Store Information</h2>
                </div>
                <div className="space-y-4">
                <div className="flex items-center space-x-4">
                  {store.logo ? (
                    <img src={store.logo} alt="Logo" className="h-12 w-12 rounded object-contain bg-white border" />
                  ) : (
                    <div className="h-12 w-12 rounded bg-gray-100 border flex items-center justify-center text-gray-400">Logo</div>
                  )}
                  <div>
                    <input
                      type="file"
                      accept="image/*"
                      onChange={async (e) => {
                        if (!e.target.files || e.target.files.length === 0) return
                        const file = e.target.files[0]
                        try {
                          const res = await api.uploadStoreLogo(parseInt(storeId), file)
                          const url = (res.data as any).url as string
                          await api.updateStore(store.id.toString(), { logo: url })
                          toast.success('Logo updated')
                          fetchStore()
                        } catch (err) {
                          toast.error('Failed to upload logo')
                        }
                      }}
                      className="block text-sm"
                    />
                  </div>
                </div>
                <div className="flex items-center space-x-4">
                  {store.favicon ? (
                    <img src={store.favicon} alt="Favicon" className="h-8 w-8 rounded object-contain bg-white border" />
                  ) : (
                    <div className="h-8 w-8 rounded bg-gray-100 border flex items-center justify-center text-gray-400 text-xs">Icon</div>
                  )}
                  <div>
                    <input
                      type="file"
                      accept="image/*"
                      onChange={async (e) => {
                        if (!e.target.files || e.target.files.length === 0) return
                        const file = e.target.files[0]
                        try {
                          const res = await api.uploadStoreFavicon(parseInt(storeId), file)
                          const url = (res.data as any).url as string
                          await api.updateStore(store.id.toString(), { favicon: url })
                          toast.success('Favicon updated')
                          fetchStore()
                        } catch (err) {
                          toast.error('Failed to upload favicon')
                        }
                      }}
                      className="block text-sm"
                    />
                  </div>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    Store Name
                  </label>
                  <p className="text-gray-900 dark:text-gray-100 break-words">{store.name}</p>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    Description
                  </label>
                  <p className="text-gray-900 dark:text-gray-100 break-words">
                    {store.description || 'No description provided'}
                  </p>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    Store URL
                  </label>
                  <p className="text-gray-900 dark:text-gray-100 break-words">
                    {store.domain || `${store.subdomain}.storemaker.com`}
                  </p>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    Created
                  </label>
                  <p className="text-gray-900 dark:text-gray-100">
                    {new Date(store.created_at).toLocaleDateString()}
                  </p>
                </div>
                </div>
              </div>
            </div>
          </div>

          <div className="space-y-6">
            {/* Quick Actions */}
            <div className="relative overflow-hidden rounded-2xl bg-white/80 dark:bg-dark-900/80 backdrop-blur-sm border border-gray-200/50 dark:border-dark-700/50 shadow-elegant">
              <div className="absolute top-0 left-0 right-0 h-24 bg-gradient-to-br from-primary-500/10 to-secondary-500/10"></div>
              <div className="relative p-6">
                <div className="flex items-center mb-5">
                  <div className="p-2 bg-gradient-to-br from-primary-500 to-secondary-500 rounded-lg mr-2.5">
                    <svg className="h-5 w-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                    </svg>
                  </div>
                  <h3 className="text-lg font-bold text-gray-900 dark:text-white">Quick Actions</h3>
                </div>
                <div className="space-y-3">
                  <Link
                    href={`/dashboard/stores/${store.id}/builder`}
                    className="group w-full flex items-center justify-between px-4 py-3.5 bg-gradient-to-r from-primary-600 to-secondary-600 hover:from-primary-700 hover:to-secondary-700 text-white rounded-xl shadow-md hover:shadow-lg hover:shadow-primary-500/50 transition-all duration-300"
                    title="Design your store's home page layout and components"
                  >
                    <div className="flex items-center">
                      <svg className="w-5 h-5 mr-3 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 4a2 2 0 114 0v1a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-1a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-1a2 2 0 10-4 0v1a1 1 0 01-1 1H7a1 1 0 01-1-1v-3a1 1 0 00-1-1H4a2 2 0 110-4h1a1 1 0 001-1V7a1 1 0 011-1h3a1 1 0 001-1V4z" />
                      </svg>
                      <span className="font-semibold">Design Home Page</span>
                    </div>
                    <svg className="w-5 h-5 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                    </svg>
                  </Link>
                  
                  <Link
                    href={`/dashboard/stores/${store.id}/pages`}
                    className="group w-full flex items-center justify-between px-4 py-3 bg-gray-100 hover:bg-gray-200 dark:bg-dark-800 dark:hover:bg-dark-700 text-gray-900 dark:text-white rounded-xl transition-all duration-200"
                    title="Manage your store pages and content"
                  >
                    <div className="flex items-center">
                      <svg className="w-5 h-5 mr-3 text-gray-600 dark:text-gray-400 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                      </svg>
                      <span className="font-medium">Manage Pages</span>
                    </div>
                    <svg className="w-5 h-5 text-gray-400 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                    </svg>
                  </Link>

                  <button 
                    onClick={handleDownloadSource}
                    disabled={isDownloading}
                    className="group w-full flex items-center justify-between px-4 py-3 bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-700 hover:to-emerald-700 disabled:from-gray-400 disabled:to-gray-500 text-white rounded-xl shadow-md hover:shadow-lg transition-all duration-200"
                    title="Download complete store source code as ZIP"
                  >
                    <div className="flex items-center">
                      <svg className="w-5 h-5 mr-3 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
                      </svg>
                      <span className="font-semibold">{isDownloading ? 'Downloading...' : 'Download Source Code'}</span>
                    </div>
                    <svg className="w-5 h-5 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                    </svg>
                  </button>
                  
                  <Link
                    href={`/dashboard/stores/${store.id}/products`}
                    className="group w-full flex items-center justify-between px-4 py-3 bg-gray-100 hover:bg-gray-200 dark:bg-dark-800 dark:hover:bg-dark-700 text-gray-900 dark:text-white rounded-xl transition-all duration-200"
                    title="Add, edit, and manage your store products"
                  >
                    <div className="flex items-center">
                      <svg className="w-5 h-5 mr-3 text-gray-600 dark:text-gray-400 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                      </svg>
                      <span className="font-medium">Manage Products</span>
                    </div>
                    <svg className="w-5 h-5 text-gray-400 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                    </svg>
                  </Link>
                  
                  <Link
                    href={`/dashboard/stores/${store.id}/orders`}
                    className="group w-full flex items-center justify-between px-4 py-3 bg-white dark:bg-dark-800 border-2 border-gray-200 dark:border-dark-700 hover:border-primary-500 dark:hover:border-primary-500 text-gray-900 dark:text-white rounded-xl transition-all duration-200 shadow-sm hover:shadow-md"
                    title="View and manage customer orders"
                  >
                    <div className="flex items-center">
                      <svg className="w-5 h-5 mr-3 text-gray-600 dark:text-gray-400 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
                      </svg>
                      <span className="font-medium">View Orders</span>
                    </div>
                    <svg className="w-5 h-5 text-gray-400 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                    </svg>
                  </Link>
                </div>
              </div>
            </div>

            {/* Stats Overview */}
            <div className="relative overflow-hidden rounded-2xl bg-white/80 dark:bg-dark-900/80 backdrop-blur-sm border border-gray-200/50 dark:border-dark-700/50 shadow-elegant">
              <div className="absolute top-0 left-0 right-0 h-24 bg-gradient-to-br from-success-500/10 to-primary-500/10"></div>
              <div className="relative p-6">
                <div className="flex items-center mb-5">
                  <div className="p-2 bg-gradient-to-br from-success-500 to-primary-500 rounded-lg mr-2.5">
                    <svg className="h-5 w-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                    </svg>
                  </div>
                  <h3 className="text-lg font-bold text-gray-900 dark:text-white">Store Stats</h3>
                </div>
                <div className="space-y-4">
                  <div className="flex items-center justify-between p-3 bg-gradient-to-r from-primary-50 to-secondary-50 dark:from-primary-900/20 dark:to-secondary-900/20 rounded-xl">
                    <div className="flex items-center">
                      <div className="p-2 bg-primary-500 rounded-lg mr-3">
                        <svg className="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                        </svg>
                      </div>
                      <span className="text-sm font-medium text-gray-700 dark:text-gray-300">Total Products</span>
                    </div>
                    <span className="text-2xl font-bold text-primary-600 dark:text-primary-400">0</span>
                  </div>
                  <div className="flex items-center justify-between p-3 bg-gradient-to-r from-success-50 to-primary-50 dark:from-success-900/20 dark:to-primary-900/20 rounded-xl">
                    <div className="flex items-center">
                      <div className="p-2 bg-success-500 rounded-lg mr-3">
                        <svg className="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
                        </svg>
                      </div>
                      <span className="text-sm font-medium text-gray-700 dark:text-gray-300">Total Orders</span>
                    </div>
                    <span className="text-2xl font-bold text-success-600 dark:text-success-400">0</span>
                  </div>
                  <div className="flex items-center justify-between p-3 bg-gradient-to-r from-secondary-50 to-accent-50 dark:from-secondary-900/20 dark:to-accent-900/20 rounded-xl">
                    <div className="flex items-center">
                      <div className="p-2 bg-secondary-500 rounded-lg mr-3">
                        <svg className="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                      </div>
                      <span className="text-sm font-medium text-gray-700 dark:text-gray-300">Revenue</span>
                    </div>
                    <span className="text-2xl font-bold text-secondary-600 dark:text-secondary-400">$0.00</span>
                  </div>
                  <div className="flex items-center justify-between p-3 bg-gradient-to-r from-accent-50 to-primary-50 dark:from-accent-900/20 dark:to-primary-900/20 rounded-xl">
                    <div className="flex items-center">
                      <div className="p-2 bg-accent-500 rounded-lg mr-3">
                        <svg className="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                        </svg>
                      </div>
                      <span className="text-sm font-medium text-gray-700 dark:text-gray-300">Visitors</span>
                    </div>
                    <span className="text-2xl font-bold text-accent-600 dark:text-accent-400">0</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </StoreLayout>
  )
}