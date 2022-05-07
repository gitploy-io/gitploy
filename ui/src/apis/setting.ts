// default authentication credentials. In product cookie-based
// authentication is being used.
export const headers = new Headers(
  process.env.REACT_APP_GITPLOY_TOKEN
    ? {
        Authorization: `Bearer ${process.env.REACT_APP_GITPLOY_TOKEN}`,
        'Content-Type': 'application/json',
      }
    : {
        'Content-Type': 'application/json',
      }
);

// default server api token.
export const token: string | undefined = process.env.REACT_APP_GITPLOY_TOKEN;

// default server address.
export const instance: string = process.env.REACT_APP_GITPLOY_SERVER || '';
