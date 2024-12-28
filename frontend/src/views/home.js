import WalletConnectionButton from "@/layout/auth/Wallet/WalletConnectionButton"
import Can from "@/layout/components/acl/Can"
import { Box, Button, Card, CardContent, Typography } from "@mui/material"

const Home = () => {
    console.log("wqeqwe", <Button>Test2</Button>);
    console.log("wqeqwe123", <Button variant="gradient">Test</Button>);

    return (
        <div>
            <h1>Home</h1>

            <Button>Test</Button>

            <br />
            <br />

            <Button color="info" variant="outlined">Test</Button>

            <br />
            <br />

            <Button color="info" variant="gradient">Test</Button>

            <br />
            <br />

            <Typography variant="h1">h1</Typography>
            <Typography variant="h2">h2</Typography>
            <Typography variant="h3">h3</Typography>
            <Typography variant="h4">h4</Typography>
            <Typography variant="h5">h5</Typography>
            <Typography variant="h6">h6</Typography>
            <Typography variant="subtitle">subtitle1</Typography>
            <Typography variant="subtitle2">subtitle2</Typography>
            <Typography variant="body">body1</Typography>
            <Typography variant="body2">body2</Typography>
            <Typography variant="caption">caption1</Typography>
            <Typography variant="caption2">caption2</Typography>
            <Typography variant="link">link1</Typography>
            <Typography variant="link2">link2</Typography>

            <br />
            <br />


            <Box sx={{ maxWidth: "300px" }}>

                <Card variant="gradient">
                    <CardContent>
                        asdsadsad
                        <br />
                        <br />
                        <br />
                        asdsadsad
                    </CardContent>

                    <div className="cardd-button"></div>
                </Card>
            </Box>
            <br />
            <br />
            <Box sx={{ maxWidth: "300px" }}>

                <Card variant="special">
                    <CardContent>
                        asdsadsad
                        <br />
                        <br />
                        <br />
                        asdsadsad
                    </CardContent>

                    <div className="cardd-button"></div>
                </Card>
            </Box>

            <br />
            <br />
            <Box sx={{ maxWidth: "300px" }}>

                <Card variant="flat">
                    <CardContent>
                        asdsadsad
                        <br />
                        <br />
                        <br />
                        asdsadsad
                    </CardContent>

                    <div className="cardd-button"></div>
                </Card>
            </Box>

            <br />
            <br />
            <Box sx={{ maxWidth: "300px" }}>

                <Card>
                    <CardContent>
                        asdsadsad
                        <br />
                        <br />
                        <br />
                        asdsadsad
                    </CardContent>

                    <div className="cardd-button"></div>
                </Card>
            </Box>


            <br />
            <br />



            <WalletConnectionButton />

            <Can I="read" a="wallet">
                If you see this message. Your wallet has been connected
            </Can>
        </div>
    )
}

export default Home