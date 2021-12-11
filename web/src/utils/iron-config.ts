const ironConfig = {
  password: process.env.NEXT_PUBLIC_COOKIE_KEY || '38434289u239832',
  cookieName: "fullcycle-session",
  cookieOptions: {
    // the next line allows to use the session in non-https environments like
    // Next.js dev mode (http://localhost:3000)
    secure: process.env.NODE_ENV === "production" ? true : false,
  },
};

declare module "iron-session" {
  interface IronSessionData {
    account?: {
      id: number;
      name: string;
      token: string;
    };
  }
}

export default ironConfig;
