import Script from 'next/script';
import Layout from '../components/Layout';
import '../styles/globals.css';

function MyApp({ Component, pageProps }) {
  return (
    <>
      <Layout>
        <Component {...pageProps} />
      </Layout>
      <Script src="/wasm_exec.js" />
    </>
  );
}

export default MyApp;
