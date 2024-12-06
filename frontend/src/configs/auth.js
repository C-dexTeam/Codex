let baseUrl = process.env.NEXT_PUBLIC_BASE_URL

const authConfig = {
  refresh: baseUrl + '/private/user/profile',
  login: baseUrl + '/public/login',
  logout: baseUrl + '/public/logout',
  register: baseUrl + '/public/register',
  wallet: baseUrl + '/public/wallet',

  session: 'userSession',

  homeRoute: {
    'First-Login': '/register/wallet',
    'nowallet-member': '/home',
    'member': '/home',
    'admin': '/admin',
  }
};

export default authConfig;
