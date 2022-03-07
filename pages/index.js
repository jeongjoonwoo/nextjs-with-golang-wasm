import { useState } from 'react';

const Page = () => {
  const [value, setValue] = useState(0);
  const onClick = () => {
    const go = new Go();

    WebAssembly.instantiateStreaming(fetch('/cal.wasm'), go.importObject).then(
      (result) => {
        go.run(result.instance);

        setValue(sum);
      }
    );
  };

  return (
    <>
      <p>{value}</p>
      <button onClick={onClick}>hello</button>
    </>
  );
};

export default Page;
