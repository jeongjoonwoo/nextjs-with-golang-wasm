import Link from "next/link"


export default function NavBar(){
   return (
      <nav>
         <div>
            <Link href="/">
               <a>Home</a>   
            </Link>
            <Link href="/wasm/client">
               <a>Client</a>   
            </Link>
            <Link href="/wasm/server">
               <a>Server</a>   
            </Link>
         </div>
         <style jsx>{`
        nav {
          display: flex;
          gap: 10px;
          flex-direction: column;
          align-items: center;
          padding-top: 20px;
          padding-bottom: 10px;
          box-shadow: rgba(50, 50, 93, 0.25) 0px 50px 100px -20px,
            rgba(0, 0, 0, 0.3) 0px 30px 60px -30px;
        }
        img {
          max-width: 100px;
          margin-bottom: 5px;
        }
        nav a {
          font-weight: 600;
          font-size: 18px;
        }
        .active {
          color: tomato;
        }
        nav div {
          display: flex;
          gap: 10px;
        }
      `}</style>
      </nav>
   )
}