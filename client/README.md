# Client-side Application

This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Project Setup

### Environment Variables

- NEXT_PUBLIC_API_ENDPOINT
- NEXT_PUBLIC_DOMAIN

For example, you can see [`.env.example`](.env.example)

### Installing Depedencies

```bash
npm install
```

### Compiles and hot-reloads for development

```bash
npm run dev
```

### Compiles and minifies for production

```bash
npm run build
```

## Dockerize The Application (Development)

1.  Build docker image
    ```bash
    docker build \
    --build-arg API_ENDPOINT="http://localhost:5000/api/v1/go" \
    --build-arg DOMAIN="http://localhost:5000" \
    -t url-client:1.0 .
    ```
2.  Run docker image
    ```bash
    docker run -d --name url-client \
    -p 3000:3000 \
    url-client:1.0
    ```

## Useful Links

To learn more about this project, take a look at the following resources:

- [React.js Documentation](https://reactjs.org/docs/getting-started.html)
- [Next.js Documentation](https://nextjs.org/docs)
- [Axios, Promise based HTTP client for the browser and node.js](https://github.com/axios/axios)
- [Getting Started with React Hook Form](https://react-hook-form.com/get-started)
- [Yup, a JavaScript schema builder for value parsing and validation](https://github.com/jquense/yup)
