import LevelBar from "@/components/bar/LevelBar"
import { SchoolOutlined } from "@mui/icons-material"
import { Box, Button, Card, CardActions, CardContent, Grid, Typography, useMediaQuery } from "@mui/material"
import { useRouter } from "next/router"
import { useEffect } from "react"

const Chapters = () => {

    const router = useRouter()

    useEffect(() => {
        console.log(router.query.course)
    }, [router.isReady])

    const _md = useMediaQuery(theme => theme.breakpoints.down('lgPlus'))

    return (
        <Grid container spacing={6}>
            <Grid item xs={12} sx={{ pt: "0px !important" }}>
                <Box
                    sx={{
                        display: "flex",
                        justifyContent: "space-between",
                        alignItems: "center",
                        height: _md ? "calc(100vh - 8.5rem)" : "auto"
                    }}
                >
                    <Box sx={{ display: "flex", flexDirection: "column", height: "100%", gap: "5rem", justifyContent: "center" }}>
                        <Box sx={{ display: "flex", flexDirection: "column", gap: "1rem" }}>
                            <Typography variant="h2" color={"primary"}>
                                <Typography variant="strong">Learn Solidity:</Typography> <Typography variant="slim">Introduction to Solidity</Typography>
                            </Typography>

                            <Typography variant="h4">
                                <Typography variant="slim">
                                    Learn about the fast-growing blockchain space and its software world.
                                </Typography>
                            </Typography>

                            <Box sx={{ display: "flex", gap: "0.5rem", alignItems: "center" }}>
                                <SchoolOutlined />

                                <Typography variant="subtitle1" sx={{ gap: "0.25rem" }}>
                                    <Typography variant="body1" color={"secondary"}>For</Typography>
                                    Beginers
                                </Typography>
                            </Box>

                        </Box>

                        <Box sx={{
                            maxWidth: "36rem",
                        }}>
                            <LevelBar
                                variant="astronout"
                                proggress={32}
                                level={3}
                                free
                            />
                        </Box>
                    </Box>

                    <Box
                        sx={{
                            position: "relative",
                            display: "inline-flex",
                            gap: "0.25rem",
                        }}
                    >
                        <Box
                            sx={{
                                position: "relative",
                                width: "12rem",
                                height: "16rem",
                                overflow: "hidden",
                                display: "inline-block",
                                justifyContent: "center",
                                alignItems: "center",
                                backgroundColor: theme => `${theme.palette.background.default}`,
                                boxShadow: "0 0 1rem 0 rgba(0, 0, 0, 0.1)",
                                "& img": {
                                    maxWidth: "calc(100%)",
                                    maxHeight: "100%",
                                    width: "calc(100%)",
                                    height: "auto",
                                    objectFit: "cover",
                                }
                            }}
                        >
                            <img src="https://picsum.photos/600/1200" />
                        </Box>

                        {/* Fotoğrafın sağ tarafına 90 derece döndürülmüş yazı yaz */}
                        <Box
                            sx={{
                                writingMode: "vertical-rl",
                                transform: "scale(-1, -1)",
                            }}
                        >
                            <Typography
                                variant="h4"
                                sx={{
                                    color: theme => `${theme.palette.text.primary}`,
                                    lineHeight: "1.5rem",
                                }}
                            >
                                End of chapter reward
                            </Typography>
                        </Box>
                    </Box>
                </Box>
            </Grid>

            <Grid item xs={12}>
                <Typography variant="h3" color="success">Chapter</Typography>
            </Grid>

            <Grid item xs={12}>
                <Box
                    sx={{
                        display: "flex",
                        width: '100%',
                        gap: "1rem",
                        height: "100%"
                    }}
                >
                    <Box
                        sx={{
                            display: "flex",
                            flexDirection: "column",
                            gap: "1rem",
                        }}
                    >
                        <Box sx={{ display: "flex", justifyContent: "center", alignItems: "center", width: '104px', height: "104px", borderRadius: "1rem", border: theme => "3px solid " + theme.palette.warning.main }}>
                            <Typography variant="h2">1</Typography>
                        </Box>
                        <Box sx={{
                            position: "relative", display: "flex", justifyContent: "center", alignItems: "center",
                            width: '110px', height: "110px", borderRadius: "1rem",
                            background: theme => `linear-gradient(to bottom right, ${theme.palette.error.main} 0%, ${theme.palette.info.main} 60%)`,
                            "& img": {
                                maxWidth: "calc(100% - 6px)",
                                maxHeight: "calc(100% - 6px)",
                                width: "100%",
                                height: "auto",
                                objectFit: "cover",
                                borderRadius: "1rem",
                            }
                        }}>
                            <img src="https://picsum.photos/200/300" />
                        </Box>
                    </Box>

                    <Card sx={{ width: "100%", display: "flex", flexDirection: "column", gap: "1rem", justifyContent: "space-between", height: "100%" }}>
                        <CardContent size="small">
                            <Typography variant="h3">First steps in Solidity</Typography>

                            <Typography variant="caption1">Let's get started on the first steps of creating Smart Contracts with Solidity. In this class you will learn how to:</Typography>

                        </CardContent>

                        <CardActions flat sx={{ display: "flex", justifyContent: "flex-end" }}>
                            <Button color="primary" bgc sx={{ minWidth: "8rem" }}>Start</Button>
                        </CardActions>
                    </Card>
                </Box>
            </Grid>
        </Grid>
    )
}

export default Chapters