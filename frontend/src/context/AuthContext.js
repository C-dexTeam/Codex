// ** React Imports
import { createContext, useEffect, useState } from 'react'
// ** Next Import
import { useRouter } from 'next/router'
// ** Axios
import authConfig from '@/configs/auth'
import axios from 'axios'
import { showToast } from '@/utils/showToast'
import { useWallet } from '@solana/wallet-adapter-react'
import { useWalletModal } from '@solana/wallet-adapter-react-ui'
import { SigninMessage } from '@/layout/auth/Wallet/SignInMessage'
import { binary_to_base58 } from 'base58-js'
import { t } from 'i18next'

// ** Defaults
const defaultProvider = {
  user: {
    id: null,
    username: "user",
    email: null,
    name: "Space",
    surname: "Hunter",
    roleID: null,
    role: "public",
  },
  loading: true,
  setUser: () => null,
  setLoading: () => Boolean,
  login: () => Promise.resolve(),
  logout: () => Promise.resolve(),
  register: () => Promise.resolve(),
  refreshAuth: () => Promise.resolve(),
}

const AuthContext = createContext(defaultProvider)

const AuthProvider = ({ children }) => {
  // ** States
  const [user, setUser] = useState(defaultProvider.user)
  const [loading, setLoading] = useState(defaultProvider.loading)

  // ** Hooks
  const router = useRouter()

  // ** Wallet
  const wallet = useWallet();
  const { disconnect } = useWallet();
  const walletModal = useWalletModal();

  const createSessionData = (data) => {
    const userData = {
      id: data?.userID || defaultProvider.id,
      username: data?.username || defaultProvider.username,
      email: data?.email || defaultProvider.email,
      name: data?.name || defaultProvider.name,
      surname: data?.surname || defaultProvider.surname,
      roleID: data?.roleID || defaultProvider.roleID,
      role: data?.roleName || data?.role || defaultProvider.role,
    }

    return userData
  }

  const createSession = (data) => {
    const userData = createSessionData(data)

    setUser(userData) // Set the user data to the state
    localStorage.setItem(authConfig.session, JSON.stringify(userData)) // Set the user data to the local storage
  }

  const checkSession = (realData = null) => { // returns boolean
    const session = localStorage.getItem(authConfig.session)

    if (!session) return false

    if (realData) {
      const sessionData = JSON.parse(session)
      const realDataFormatted = createSessionData(realData)

      const isEqual = Object.entries(sessionData).every(([key, value]) => realDataFormatted[key] === value);

      if (!isEqual) return false;
    }

    return true
  }

  const deleteStorage = () => {
    setUser(null)
    setLoading(false)
    localStorage.removeItem(authConfig.session)

    // const firstPath = router.pathname.split('/')[1]
    // if (firstPath != 'login') window.location.href = '/login'
  }

  const restoreStorage = () => {
    setUser(defaultProvider.user)
    setLoading(false)
    localStorage.setItem(authConfig.session, JSON.stringify(defaultProvider.user))
  }

  const openSignIn = () => {
    console.log("wallet.connected", wallet.connected);

    if (wallet.connected) handleSignIn();
  }

  /** Handle User Login Function
   * 
   * @param {String} username
   * @param {String} password
   */
  const handleLogin = async (data) => {
    // ** Set loading to true
    setLoading(true)

    // Send a POST request to the API
    axios.post(authConfig.login, {
      username: data.username || null,
      password: data.password || null,
    })
      .then(response => {
        // ** If the status code is 200, It means the login is successful
        if (response.data?.statusCode === 200) {
          // ** Send a success message to the user
          showToast("dismiss") // Dismiss the previous toast if it exists
          showToast("success", response.data?.message) // Show the success message

          createSession(response.data?.data) // Create a session for the user

          setLoading(false) // Set loading to false

          router.push("/") // Redirect the user to the home page
        } else {
          // ** If the status code is not 200, It means the login is not successful
          // ** Send an error message to the user
          showToast("dismiss") // Dismiss the previous toast if it exists
          showToast("error", response.data?.message) // Show the error message

          deleteStorage() // Delete the user data
        }

      })
      .catch(error => {
        // ** If an error occurs, Send an error message to the user
        showToast("dismiss") // Dismiss the previous toast if it exists
        showToast("error", error.response.data.message) // Show the error message
        console.error(error) // Log the error to the console
        deleteStorage() // Delete the user data
      })
  }

  /**
   * 
   * @param {String} message
   * @param {String} publicKeyBase58
   * @param {String} signatureBase58
   */
  const handleWalletLogin = async (data) => {
    try {
      const response = await axios({
        url: authConfig.wallet,
        method: "POST",
        data: data,
      });

      if (response.status === 200) {
        const user = response?.data?.data;
        createSession(user);

        showToast("dismiss");
        showToast("success", "Logged in successfully");
      } else {
        showToast("dismiss");
        showToast("error", response.data.message);
      }
    } catch (error) {
      showToast("dismiss");
      showToast("error", t(error?.response?.data?.message));
    }
  };

  const handleSignIn = async () => {
    try {
      // if wallet is not connected then show wallet modal
      if (!wallet.connected) walletModal.setVisible(true)

      // if wallet public key or signMessage is not available then return
      if (!wallet.publicKey || !wallet.signMessage) return;

      const pKey = localStorage.getItem(authConfig.publicKey) || null;
      // if user already logged in do not then return 
      if (user?.publicKey && pKey) return;

      // create a new SigninMessage object
      const message = new SigninMessage({
        domain: window.location.host,
        publicKey: wallet.publicKey.toBase58(),
        statement: `Sign in to ${window.location.host}`,
      });

      const data = new TextEncoder().encode(message.prepare()); // encode the message
      const signature = await wallet.signMessage(data); // sign the message
      const serializedSignature = binary_to_base58(signature); // convert the signature to base58

      // call the handleWalletLogin function with the message, public key and signature to communicate with the backend
      handleWalletLogin({
        message: message.statement, // message statement in plain text
        publicKeyBase58: wallet.publicKey, // public key in base58
        signatureBase58: serializedSignature, // signature in base58
      });
    } catch (error) {
      console.log(error);
    }
  }

  //** Handle User Logout Function
  const handleLogout = () => {
    // ** Send a POST request to the API
    axios.post(authConfig.logout)
      .then(response => {
        // ** If the status code is 200, It means the logout is successful
        if (response.data?.statusCode === 200) {
          // ** Send a success message to the user
          showToast("dismiss") // Dismiss the previous toast if it exists
          showToast("success", response.data?.message) // Show the success message

          deleteStorage() // Delete the user data
          router.push("/") // Redirect the user to the login page
        } else {
          // ** If the status code is not 200, It means the logout is not successful
          // ** Send an error message to the user
          showToast("dismiss") // Dismiss the previous toast if it exists
          showToast("error", response.data?.message) // Show the error message

          deleteStorage() // Delete the user data
          router.push("/") // Redirect the user to the login page
        }
      })
      .catch(error => {
        // ** If an error occurs, Send an error message to the user
        showToast("dismiss") // Dismiss the previous toast if it exists
        showToast("error", "An error occurred. Please try again") // Show the error message
        console.error(error) // Log the error to the console
        deleteStorage() // Delete the user data
        router.push("/") // Redirect the user to the login page
      })
  }

  /** Handle User Register Function
   * @param {String} confirmPassword
   * @param {String} email
   * @param {String} username
   * @param {String} password
   */
  const handleRegister = async (data) => {
    // ** Set loading to true
    setLoading(true)

    // Send a POST request to the API
    axios.post(authConfig.register, {
      username: data?.username || null,
      email: data?.email || null,
      password: data?.password || null,
      ConfirmPassword: data?.confirmPassword || null,
    })
      .then(response => {
        // ** If the status code is 200, It means the registration is successful
        if (response.data?.statusCode === 200) {
          // ** Send a success message to the user
          showToast("dismiss") // Dismiss the previous toast if it exists
          showToast("success", response.data?.message) // Show the success message

          // router.push("/login") // Redirect the user to the login page

          setLoading(false) // Set loading to false
        } else {
          // ** If the status code is not 200, It means the registration is not successful
          // ** Send an error message to the user
          showToast("dismiss") // Dismiss the previous toast if it exists
          showToast("error", response.data?.message) // Show the error message

          deleteStorage() // Delete the user data
        }
      })
      .catch(error => {
        // ** If an error occurs, Send an error message to the user
        showToast("dismiss") // Dismiss the previous toast if it exists
        showToast("error", "An error occurred. Please try again") // Show the error message
        console.error(error) // Log the error to the console
        deleteStorage() // Delete the user data
      })

  }

  // ** Refresh User's Auth Function
  const refreshAuth = async () => {
    setLoading(true) // ** Set loading to true

    // Send a GET request to the API
    axios.get(authConfig.refresh)
      .then(response => {
        console.log("qweqweqe", response);

        // ** If the status code is 200, It means the user is authenticated
        if (response.data?.statusCode === 200) {
          createSession(response.data?.data) // Create a session for the user

          setLoading(false) // Set loading to false
        } else {
          // ** If the status code is not 200, It means the user is not authenticated
          // ** Send an error message to the user
          showToast("dismiss") // Dismiss the previous toast if it exists
          showToast("error", response.data?.message) // Show the error message

          restoreStorage() // Restore the user data
        }
      })
      .catch(error => {
        console.error(error) // Log the error to the console
        restoreStorage() // Restore the user data
      })
  }

  useEffect(() => {
    // if (!wallet.connected && user?.publicKey) {
    //   handleLogout()
    //   return
    // }

    if (wallet.connected) openSignIn()
    else refreshAuth();
  }, [wallet]);

  const values = {
    user,
    loading,
    setUser,
    setLoading,
    login: handleLogin,
    logout: handleLogout,
    register: handleRegister,
    refresh: refreshAuth,
  }

  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>
}

export { AuthContext, AuthProvider }
