"use client"

import { useState } from 'react'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { useAuth } from '@/contexts/auth-context'
import {
  HomeIcon,
  BuildingStorefrontIcon,
  Cog6ToothIcon,
  UserIcon,
  ArrowRightOnRectangleIcon,
  Bars3Icon,
  XMarkIcon,
} from '@heroicons/react/24/outline'

interface DashboardLayoutProps {
  children: React.ReactNode
}

const navigation = [
  { name: 'My Stores', href: '/dashboard', icon: BuildingStorefrontIcon },
]

export function DashboardLayout({ children }: DashboardLayoutProps) {
  const [sidebarOpen, setSidebarOpen] = useState(false)
  const { user, logout } = useAuth()
  const router = useRouter()

  const handleLogout = () => {
    logout()
    router.push('/')
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-50 via-primary-50/30 to-secondary-50/30 dark:from-dark-950 dark:via-dark-900 dark:to-dark-950">
      {/* Mobile sidebar */}
      <div className={`fixed inset-0 z-50 lg:hidden ${sidebarOpen ? 'block' : 'hidden'}`}>
        <div className="fixed inset-0 bg-gray-900/80 backdrop-blur-sm" onClick={() => setSidebarOpen(false)} />
        
        <div className="fixed inset-y-0 left-0 flex w-72 flex-col bg-white/95 dark:bg-dark-900/95 backdrop-blur-xl shadow-2xl animate-slide-right">
          {/* Logo section with gradient */}
          <div className="flex items-center justify-between h-16 px-6 border-b border-gray-200/50 dark:border-dark-700/50 bg-gradient-to-r from-primary-500 to-secondary-500">
            <Link href="/" className="text-xl font-bold text-white flex items-center">
              <div className="w-8 h-8 rounded-lg bg-white/20 backdrop-blur-sm flex items-center justify-center mr-2">
                <BuildingStorefrontIcon className="h-5 w-5 text-white" />
              </div>
              StoreMaker
            </Link>
            <button
              onClick={() => setSidebarOpen(false)}
              className="text-white/80 hover:text-white p-2 rounded-lg hover:bg-white/10 transition-all"
            >
              <XMarkIcon className="h-6 w-6" />
            </button>
          </div>
          
          <nav className="flex-1 px-4 py-6 space-y-2 overflow-y-auto">
            {navigation.map((item) => (
              <Link
                key={item.name}
                href={item.href}
                className="flex items-center px-4 py-3 text-sm font-medium rounded-xl text-gray-700 dark:text-gray-300 hover:bg-gradient-to-r hover:from-primary-50 hover:to-secondary-50 dark:hover:from-primary-900/20 dark:hover:to-secondary-900/20 hover:text-primary-700 dark:hover:text-primary-400 transition-all duration-200 group"
                onClick={() => setSidebarOpen(false)}
              >
                <item.icon className="h-5 w-5 mr-3 group-hover:scale-110 transition-transform" />
                {item.name}
              </Link>
            ))}
          </nav>
          
          <div className="px-4 py-4 border-t border-gray-200/50 dark:border-dark-700/50 bg-gray-50/50 dark:bg-dark-800/50">
            <button
              onClick={handleLogout}
              className="flex items-center w-full px-4 py-3 text-sm font-medium text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-xl transition-all duration-200 group"
            >
              <ArrowRightOnRectangleIcon className="h-5 w-5 mr-3 group-hover:scale-110 transition-transform" />
              Sign Out
            </button>
          </div>
        </div>
      </div>

      {/* Desktop sidebar */}
      <div className="hidden lg:fixed lg:inset-y-0 lg:flex lg:w-72 lg:flex-col z-40">
        <div className="flex flex-col flex-grow bg-white/80 dark:bg-dark-900/80 backdrop-blur-xl shadow-elegant border-r border-gray-200/50 dark:border-dark-700/50">
          {/* Logo section with gradient */}
          <div className="relative h-20 px-6 flex items-center overflow-hidden">
            <div className="absolute inset-0 bg-gradient-to-r from-primary-500 to-secondary-500"></div>
            <div className="absolute inset-0 bg-gradient-to-br from-transparent via-white/10 to-white/20"></div>
            <Link href="/" className="relative text-2xl font-bold text-white flex items-center group">
              <div className="w-10 h-10 rounded-xl bg-white/20 backdrop-blur-sm flex items-center justify-center mr-3 group-hover:scale-110 transition-transform duration-300">
                <BuildingStorefrontIcon className="h-6 w-6 text-white" />
              </div>
              <span className="bg-clip-text">StoreMaker</span>
            </Link>
          </div>
          
          <nav className="flex-1 px-4 py-6 space-y-2 overflow-y-auto sidebar-scrollbar">
            {navigation.map((item) => (
              <Link
                key={item.name}
                href={item.href}
                className="flex items-center px-4 py-3.5 text-sm font-semibold rounded-xl text-gray-700 dark:text-gray-300 hover:bg-gradient-to-r hover:from-primary-50 hover:to-secondary-50 dark:hover:from-primary-900/20 dark:hover:to-secondary-900/20 hover:text-primary-700 dark:hover:text-primary-400 hover:shadow-sm transition-all duration-200 group"
              >
                <item.icon className="h-5 w-5 mr-3 group-hover:scale-110 transition-transform" />
                {item.name}
              </Link>
            ))}
          </nav>
          
          <div className="px-4 py-6 border-t border-gray-200/50 dark:border-dark-700/50 bg-gradient-to-br from-gray-50/50 to-primary-50/30 dark:from-dark-800/50 dark:to-primary-900/10">
            <div className="flex items-center mb-4 px-3 py-3 bg-white/50 dark:bg-dark-800/50 rounded-xl backdrop-blur-sm">
              <div className="relative">
                <div className="w-10 h-10 bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl flex items-center justify-center shadow-lg">
                  <span className="text-white text-sm font-bold">
                    {user?.first_name?.charAt(0) || user?.email?.charAt(0)}
                  </span>
                </div>
                <div className="absolute -bottom-1 -right-1 w-4 h-4 bg-success-500 border-2 border-white dark:border-dark-800 rounded-full"></div>
              </div>
              <div className="ml-3 min-w-0 flex-1">
                <p className="text-sm font-bold text-gray-900 dark:text-white truncate">
                  {user?.first_name ? `${user.first_name} ${user.last_name}` : user?.email}
                </p>
                <p className="text-xs font-medium text-gray-600 dark:text-gray-400 capitalize truncate flex items-center">
                  <span className="w-2 h-2 rounded-full bg-primary-500 mr-1.5 animate-pulse"></span>
                  {user?.role}
                </p>
              </div>
            </div>
            <button
              onClick={handleLogout}
              className="flex items-center w-full px-4 py-3 text-sm font-semibold text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-xl transition-all duration-200 group hover:shadow-sm"
            >
              <ArrowRightOnRectangleIcon className="h-5 w-5 mr-3 group-hover:scale-110 transition-transform" />
              Sign Out
            </button>
          </div>
        </div>
      </div>

      {/* Main content */}
      <div className="lg:pl-72">
        {/* Top navigation with glassmorphism */}
        <div className="sticky top-0 z-30 bg-white/70 dark:bg-dark-900/70 backdrop-blur-xl shadow-sm border-b border-gray-200/50 dark:border-dark-700/50">
          <div className="flex items-center justify-between h-16 px-6">
            <button
              onClick={() => setSidebarOpen(true)}
              className="lg:hidden text-gray-600 hover:text-primary-600 dark:text-gray-400 dark:hover:text-primary-400 p-2 rounded-xl hover:bg-primary-50 dark:hover:bg-primary-900/20 transition-all"
            >
              <Bars3Icon className="h-6 w-6" />
            </button>
            
            <div className="flex items-center space-x-3">
              <div className="hidden sm:flex items-center px-4 py-2 bg-gradient-to-r from-primary-50 to-secondary-50 dark:from-primary-900/20 dark:to-secondary-900/20 rounded-xl">
                <svg className="w-5 h-5 text-primary-600 dark:text-primary-400 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z" />
                </svg>
                <span className="text-sm font-semibold bg-gradient-to-r from-primary-600 to-secondary-600 bg-clip-text text-transparent">
                  Welcome back, {user?.first_name || 'User'}!
                </span>
              </div>
              <div className="flex sm:hidden items-center px-3 py-2 bg-gradient-to-r from-primary-50 to-secondary-50 dark:from-primary-900/20 dark:to-secondary-900/20 rounded-xl">
                <span className="text-sm font-semibold text-primary-600 dark:text-primary-400">
                  Welcome!
                </span>
              </div>
            </div>
          </div>
        </div>

        {/* Page content with animation */}
        <main className="p-6 animate-fade-in">
          {children}
        </main>
      </div>
    </div>
  )
}