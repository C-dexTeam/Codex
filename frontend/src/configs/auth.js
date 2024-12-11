let baseUrl = process.env.NEXT_PUBLIC_BASE_URL

const authConfig = {
  refresh: baseUrl + '/private/user/profile',
  login: baseUrl + '/public/login',
  logout: baseUrl + '/public/logout',
  register: baseUrl + '/public/register',
  wallet: baseUrl + '/public/wallet',

  publicKey: 'publicKey',
  session: 'userSession',

  homeRoute: {
    'First-Login': '/register/wallet',
    'public': '/',
    'member': '/',
    'admin': '/admin',
  }
};

export default authConfig;
