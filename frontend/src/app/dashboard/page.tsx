"use client"

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { useAuth } from '@/contexts/auth-context'
import { api } from '@/lib/api'
import { Store } from '@/types'
import { DashboardLayout } from '@/components/layout/dashboard-layout'
import { StoreCard } from '@/components/dashboard/store-card'
import { CreateStoreModal } from '@/components/dashboard/create-store-modal'
import { PlusIcon, BuildingStorefrontIcon } from '@heroicons/react/24/outline'
import toast from 'react-hot-toast'

export default function DashboardPage() {
  const { isLoading, isAuthenticated } = useAuth()
  const router = useRouter()
  const [stores, setStores] = useState<Store[]>([])
  const [isLoadingStores, setIsLoadingStores] = useState(true)
  const [showCreateModal, setShowCreateModal] = useState(false)

  useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      router.push('/auth/login')
      return
    }

    if (isAuthenticated) {
      fetchStores()
    }
  }, [isLoading, isAuthenticated, router])

  const fetchStores = async () => {
    try {
      setIsLoadingStores(true)
      const response = await api.getStores()
      console.log('Stores response:', response.data) // Debug log
      
      // Handle different response formats
      const data = response.data as any
      let storesData = []
      if (Array.isArray(data)) {
        storesData = data
      } else if (data && typeof data === 'object' && 'data' in data && Array.isArray(data.data)) {
        storesData = data.data
      } else if (data && typeof data === 'object' && 'stores' in data && Array.isArray(data.stores)) {
        storesData = data.stores
      }
      
      console.log('Parsed stores data:', storesData) // Debug log
      setStores(storesData)
    } catch (error) {
      console.error('Failed to fetch stores:', error)
      toast.error('Failed to load stores')
      setStores([]) // Ensure stores is always an array
    } finally {
      setIsLoadingStores(false)
    }
  }

  const handleStoreCreated = (newStore: Store) => {
    setStores(prev => Array.isArray(prev) ? [...prev, newStore] : [newStore])
    setShowCreateModal(false)
    toast.success('Store created successfully!')
    // Refresh the stores list to get the complete data
    setTimeout(() => {
      fetchStores()
    }, 1000)
  }

  const handleDeleteStore = async (storeId: number) => {
    if (!storeId || !confirm('Are you sure you want to delete this store?')) return

    try {
      await api.deleteStore(storeId)
      setStores(prev => Array.isArray(prev) ? prev.filter(store => store.id !== storeId) : [])
      toast.success('Store deleted successfully')
    } catch {
      toast.error('Failed to delete store')
    }
  }

  if (isLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="spinner h-8 w-8"></div>
      </div>
    )
  }

  if (!isAuthenticated) {
    return null // Will redirect
  }

  return (
    <DashboardLayout>
      <div className="space-y-6 sm:space-y-8">
        {/* Header */}
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between space-y-4 sm:space-y-0">
          <div className="space-y-2">
            <h1 className="text-2xl sm:text-3xl font-bold text-gray-900 dark:text-gray-100">
              My Stores
            </h1>
            <p className="text-sm sm:text-base text-gray-600 dark:text-gray-400">
              Manage your ecommerce stores and track their performance
            </p>
          </div>
          <button
            onClick={() => setShowCreateModal(true)}
            className="btn-primary flex items-center justify-center sm:justify-start w-full sm:w-auto"
          >
            <PlusIcon className="h-5 w-5 mr-2" />
            Create Store
          </button>
        </div>

        {/* Stats Overview with Gradient Cards */}
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 animate-slide-up">
          {/* Total Stores Card */}
          <div className="group relative overflow-hidden rounded-2xl transition-all duration-300 hover:scale-105">
            <div className="absolute inset-0 bg-gradient-to-br from-primary-500 to-primary-700 opacity-90 group-hover:opacity-100 transition-opacity"></div>
            <div className="absolute inset-0 bg-gradient-to-br from-transparent via-white/10 to-white/20"></div>
            <div className="relative p-6">
              <div className="flex items-center justify-between mb-4">
                <div className="p-3 bg-white/20 backdrop-blur-sm rounded-xl">
                  <BuildingStorefrontIcon className="h-7 w-7 text-white" />
                </div>
                <div className="px-3 py-1 bg-white/20 backdrop-blur-sm rounded-full text-xs font-semibold text-white">
                  All
                </div>
              </div>
              <p className="text-white/80 text-sm font-medium mb-1">Total Stores</p>
              <p className="text-4xl font-bold text-white">{Array.isArray(stores) ? stores.length : 0}</p>
              <div className="mt-4 flex items-center text-white/70 text-xs">
                <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
                </svg>
                Your online businesses
              </div>
            </div>
          </div>
          
          {/* Active Stores Card */}
          <div className="group relative overflow-hidden rounded-2xl transition-all duration-300 hover:scale-105">
            <div className="absolute inset-0 bg-gradient-to-br from-success-500 to-success-700 opacity-90 group-hover:opacity-100 transition-opacity"></div>
            <div className="absolute inset-0 bg-gradient-to-br from-transparent via-white/10 to-white/20"></div>
            <div className="relative p-6">
              <div className="flex items-center justify-between mb-4">
                <div className="p-3 bg-white/20 backdrop-blur-sm rounded-xl">
                  <svg className="h-7 w-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div className="px-3 py-1 bg-white/20 backdrop-blur-sm rounded-full text-xs font-semibold text-white flex items-center">
                  <span className="w-2 h-2 rounded-full bg-white mr-1.5 animate-pulse"></span>
                  Live
                </div>
              </div>
              <p className="text-white/80 text-sm font-medium mb-1">Active Stores</p>
              <p className="text-4xl font-bold text-white">
                {Array.isArray(stores) ? stores.filter(store => store.status === 'active').length : 0}
              </p>
              <div className="mt-4 flex items-center text-white/70 text-xs">
                <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
                Publicly accessible
              </div>
            </div>
          </div>
          
          {/* Draft Stores Card */}
          <div className="group relative overflow-hidden rounded-2xl transition-all duration-300 hover:scale-105">
            <div className="absolute inset-0 bg-gradient-to-br from-accent-500 to-accent-700 opacity-90 group-hover:opacity-100 transition-opacity"></div>
            <div className="absolute inset-0 bg-gradient-to-br from-transparent via-white/10 to-white/20"></div>
            <div className="relative p-6">
              <div className="flex items-center justify-between mb-4">
                <div className="p-3 bg-white/20 backdrop-blur-sm rounded-xl">
                  <svg className="h-7 w-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </div>
                <div className="px-3 py-1 bg-white/20 backdrop-blur-sm rounded-full text-xs font-semibold text-white">
                  Draft
                </div>
              </div>
              <p className="text-white/80 text-sm font-medium mb-1">Draft Stores</p>
              <p className="text-4xl font-bold text-white">
                {Array.isArray(stores) ? stores.filter(store => store.status === 'draft').length : 0}
              </p>
              <div className="mt-4 flex items-center text-white/70 text-xs">
                <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
                In development
              </div>
            </div>
          </div>
          
          {/* Revenue Card */}
          <div className="group relative overflow-hidden rounded-2xl transition-all duration-300 hover:scale-105">
            <div className="absolute inset-0 bg-gradient-to-br from-secondary-500 to-secondary-700 opacity-90 group-hover:opacity-100 transition-opacity"></div>
            <div className="absolute inset-0 bg-gradient-to-br from-transparent via-white/10 to-white/20"></div>
            <div className="relative p-6">
              <div className="flex items-center justify-between mb-4">
                <div className="p-3 bg-white/20 backdrop-blur-sm rounded-xl">
                  <svg className="h-7 w-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div className="px-3 py-1 bg-white/20 backdrop-blur-sm rounded-full text-xs font-semibold text-white">
                  Month
                </div>
              </div>
              <p className="text-white/80 text-sm font-medium mb-1">This Month</p>
              <p className="text-4xl font-bold text-white">$0</p>
              <div className="mt-4 flex items-center text-white/70 text-xs">
                <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
                Total earnings
              </div>
            </div>
          </div>
        </div>

        {/* Stores Grid */}
        <div>
          {isLoadingStores ? (
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 sm:gap-6">
              {[1, 2, 3].map(i => (
                <div key={i} className="card p-4 sm:p-6 animate-pulse">
                  <div className="h-4 bg-gray-300 dark:bg-gray-700 rounded mb-4"></div>
                  <div className="h-3 bg-gray-300 dark:bg-gray-700 rounded mb-2"></div>
                  <div className="h-3 bg-gray-300 dark:bg-gray-700 rounded w-2/3"></div>
                </div>
              ))}
            </div>
          ) : !Array.isArray(stores) || stores.length === 0 ? (
            <div className="text-center py-8 sm:py-12 px-4">
              <div className="w-20 h-20 sm:w-24 sm:h-24 bg-gray-100 dark:bg-gray-800 rounded-full flex items-center justify-center mx-auto mb-4">
                <PlusIcon className="h-10 w-10 sm:h-12 sm:w-12 text-gray-400" />
              </div>
              <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-2">
                No stores yet
              </h3>
              <p className="text-sm sm:text-base text-gray-600 dark:text-gray-400 mb-6 max-w-md mx-auto">
                Create your first store to start selling online
              </p>
              <button
                onClick={() => setShowCreateModal(true)}
                className="btn-primary w-full sm:w-auto"
              >
                Create Your First Store
              </button>
            </div>
          ) : (
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 sm:gap-6">
              {Array.isArray(stores) && stores.map((store, index) => (
                <StoreCard
                  key={store.id || `store-${index}`}
                  store={store}
                  onDelete={() => handleDeleteStore(store.id)}
                />
              ))}
            </div>
          )}
        </div>
      </div>

      <CreateStoreModal
        isOpen={showCreateModal}
        onClose={() => setShowCreateModal(false)}
        onStoreCreated={handleStoreCreated}
      />
    </DashboardLayout>
  )
}