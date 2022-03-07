import styles from "../../styles/Home.module.css";

const target = `https://www.w3.org/People/mimasa/test/`

export default function ServerWasm(){
   return (
      <div className={styles.container}>
         <h1 className={styles.title}>
         Server Page
         </h1>
         <h3>
            Hello
         </h3>
         <div className={styles.title}>
            <input type="text"></input>
            <button>Click</button>
         </div>
         
      </div>
   )
}

export async function getStaticProps() {
   await import("../../public/wasm_exec.js");
 
   const go = new Go();
   const mod = await WebAssembly.compile(fs.readFileSync("public/main.wasm"));
   let inst = await WebAssembly.instantiate(mod, go.importObject);
   go.run(inst);
   
   // console.log(sum(3, 4));
   // printMessage("JS calling Go and back again!", (err, message) => {
   //   if (err) {
   //     console.error(err);
   //     return;
   //   }
 
   //   console.log(message);
   // });

   searchUrl(target, (err,result) =>{
      if (err) {
         console.error(err)
         return nil
      }
      console.log(result)
   })
 
   return { props: {} };
 }