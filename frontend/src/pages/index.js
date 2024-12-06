// ** React Imports
import { useEffect } from 'react'

// ** Layout Import
import BlankLayout from '@/layout/BlankLayout'

// ** Next Imports
import { useRouter } from 'next/router'

// ** Spinner Import
import Spinner from '@/components/spinner'

// ** Hook Imports
import { useAuth } from '@/hooks/useAuth'
import authConfig from '@/configs/auth'

export const getHomeRoute = role => {
  const homeRoute = authConfig.homeRoute[role]
  return homeRoute
}

const Home = () => {
  // ** Hooks
  const auth = useAuth()
  const router = useRouter()

  useEffect(() => {
    if (!router.isReady) {
      return
    }

    if (auth.user && auth.user.role) {
      const homeRoute = getHomeRoute(auth.user.role)

      // Redirect user to Home URL
      router.replace(homeRoute)
    } else {
      // Redirect user to Login URL
      router.replace('/login')
    }
  }, [])

  return <Spinner />
}

Home.getLayout = page => <BlankLayout>{page}</BlankLayout>

export default Home
