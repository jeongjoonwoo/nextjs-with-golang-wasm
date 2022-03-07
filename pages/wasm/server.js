import styles from '../../styles/Home.module.css';

const target = `https://www.w3.org/People/mimasa/test/`;

export default function ServerWasm() {
  const onClick = async () => {
    const go = new Go();

    WebAssembly.instantiateStreaming(fetch('/main.wasm'), go.importObject).then(
      (result) => {
        go.run(result.instance);
        const a = searchUrl(target);
        console.log(a);
      }
    );
  };

  return (
    <>
      <div className={styles.container}>
        <h1 className={styles.title}>Server Page</h1>
        <h3>Hello</h3>
        <div className={styles.title}>
          <input type="text"></input>
          <button onClick={onClick}>Click</button>
        </div>
      </div>
    </>
  );
}
