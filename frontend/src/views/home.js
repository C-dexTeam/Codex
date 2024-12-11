
// import Can from "@/layout/components/acl/Can"

import WalletConnectionButton from "@/layout/auth/Wallet/WalletConnectionButton"

const Home = () => {

    return (
        <div>
            <h1>Home</h1>

            <WalletConnectionButton />

            {/* <Can I="read" a="wallet">
                If you see this message. Your wallet has been connected
            </Can> */}
        </div>
    )
}

export default Home