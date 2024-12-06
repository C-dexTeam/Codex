import Can from "@/layout/components/acl/Can"

const Home = () => {
    return (
        <div>
            <h1>Home</h1>


            <Can I="read" a="wallet">
                If you see this message. Your wallet has been connected
            </Can>

        </div>
    )
}

export default Home