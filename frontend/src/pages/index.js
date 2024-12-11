
import authConfig from '@/configs/auth'
import HomePage from "@/views/home"

export const getHomeRoute = role => {
  const homeRoute = authConfig.homeRoute[role]
  return homeRoute
}

const Home = () => {
  // ** Hooks
  // const auth = useAuth()
  // const router = useRouter()

  // useEffect(() => {
  //   if (!router.isReady) {
  //     return
  //   }

  //   if (auth.user && auth.user.role) {
  //     const homeRoute = getHomeRoute(auth.user.role)

  //     // Redirect user to Home URL
  //     router.replace(homeRoute)
  //   } else {
  //     // Redirect user to Login URL
  //     router.replace('/login')
  //   }
  // }, [])

  return (<HomePage />)
}

// Home.getLayout = page => <BlankLayout>{page}</BlankLayout>

// Home.authGuard = false
Home.acl = {
  action: 'read',
  permission: 'home'
}
export default Home
