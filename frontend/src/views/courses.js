import LevelBar from "@/components/bar/LevelBar"
import GradientCard from "@/components/card/GradientCard"
import DefaultTextField from "@/components/form/components/DefaultTextField"
import { Search, Shield, Web } from "@mui/icons-material"
import { Box, Card, CardContent, Divider, Grid, IconButton, InputAdornment, ListItemButton, ListItemIcon, ListItemText, Tab, Tabs, Tooltip, Typography } from "@mui/material"
import Image from "next/image"
import { Fragment, useState } from "react"

const popularCourses = [
    {
        id: 1,
        image: "/images/card1.jpg",
        title: "Wormhole Tooling",
        description: "Regardless of which network development environment you are using.",
        chapters: 8,
    },
    {
        id: 2,
        image: "/images/card2.jpg",
        title: "Wormhole Tooling",
        description: "Regardless of which network development environment you are using.",
        chapters: 2,
    },
    {
        id: 3,
        image: "/images/card3.jpg",
        title: "Phantom SDK",
        description: "Regardless of which network development environment you are using.",
        chapters: 12,
    },
]

const statistics = [
    {
        id: 1,
        title: "Finished",
        value: 12,
        img: "/images/finished.png",
    },
    {
        id: 2,
        title: "Strike",
        value: 5,
        img: "/images/fire.png",
    }
]

const navigationData = [
    {
        icon: <Shield />,
        label: "Security",
        path: "/security",
    },
    {
        icon: <Web />,
        label: "Best Courses",
        path: "/best-courses",
    }
]

const Courses = () => {
    const [value, setValue] = useState(0);
    const [copy, setCopy] = useState(false);

    const copyToClipboard = (text) => {
        navigator.clipboard.writeText(text)
        setCopy(true)

        setTimeout(() => {
            setCopy(false)
        }, 3000)
    }

    return (
        <Grid container spacing={6}>
            <Grid item container xs={12} md={8} spacing={6}>
                <Grid item xs={12}>
                    <Box
                        sx={{
                            width: "400px",
                            height: "200px",
                            backgroundColor: theme => `${theme.palette.background.paper}`,
                            borderRadius: "300px",
                            filter: "blur(100px)",
                            position: "absolute",
                            zIndex: 0,
                            left: "20%",
                            top: "0%",
                        }}
                    ></Box>

                    <Card variant="special" sx={{ display: "inline" }}>
                        <Box sx={{
                            backgroundColor: theme => `${theme.palette.background.default}`,
                            width: "calc(100% - 4px)",
                            height: "calc(100% - 4px)",
                            position: "absolute",
                            left: "4px",
                            top: "4px",
                            zIndex: 0,
                            display: "flex",
                        }}></Box>

                        <CardContent sx={{ position: "relative", zIndex: 1 }}>
                            <Typography variant="h4" sx={{ justifyContent: "center", width: '100%' }}>Don't you know how to learn solana?</Typography>
                            <Typography variant="h4" sx={{ justifyContent: "center", width: '100%', mt: 1 }}>go to the <Typography variant="link" color="primary">roadmap</Typography> page and explore courses</Typography>
                        </CardContent>

                        <Image
                            src="/images/astronout/astronout-sitting.png"
                            alt="Astronout sitting"
                            width={80}
                            height={80}
                            style={{
                                position: "absolute",
                                zIndex: 99,
                                right: 0,
                                bottom: "-1.125rem",
                                width: "auto",
                                height: "124px"
                            }}
                        />
                    </Card>
                </Grid>

                <Grid item xs={12}>
                    <Typography variant="h3" color="warning">Popular Courses</Typography>
                </Grid>

                <Grid item container sx={12} spacing={4}>
                    {
                        popularCourses.map(course => (
                            <Grid item xs={12} md={4}>
                                <GradientCard
                                    description={
                                        <Fragment>
                                            <Typography variant="body" color={"primary"}>
                                                <Typography variant="body1">{course.chapters}</Typography>
                                                Chapters
                                            </Typography>
                                        </Fragment>
                                    }
                                    btnText="Explore"
                                    dpos="center"
                                >
                                    <CardContent>
                                        <Box className="CardImage">
                                            <Image
                                                src={course.image}
                                                alt={course.title}
                                                width={80}
                                                height={80}
                                                style={{
                                                    width: "100%",
                                                    height: "auto",
                                                    maxHeight: "120px",
                                                    objectFit: "cover",
                                                    borderRadius: "1rem"
                                                }}
                                            />
                                        </Box>

                                        <Typography variant="h5" sx={{ mt: 2 }}>{course.title}</Typography>
                                        <Typography variant="body" sx={{ mt: 1 }}>{course.description}</Typography>
                                    </CardContent>
                                </GradientCard>
                            </Grid>
                        ))
                    }
                </Grid>

                <Grid item xs={12}>
                    <Typography variant="h3" color="warning">All Courses</Typography>
                </Grid>

                <Grid item xs={12}>
                    <Box sx={{ display: "flex", gap: "1rem", justifyContent: "start", alignItems: "center", textAlign: "start" }}>
                        <DefaultTextField
                            noControl
                            variant="outlined"
                            placeholder="Search for courses"
                            InputProps={{
                                endAdornment: (
                                    <InputAdornment position="end">
                                        <IconButton size="small">
                                            <Search />
                                        </IconButton>
                                    </InputAdornment>
                                ),
                            }}
                        />

                        <Tabs
                            value={value}
                            onChange={(event, newValue) => setValue(newValue)}
                            scrollButtons
                            allowScrollButtonsMobile
                            variant="scrollable"
                            aria-label="Categories"
                            type="box"
                        >
                            <Tab label="Solana 101" />
                            <Tab label="JS" />
                            <Tab label="Rust" />
                            <Tab label="TS" />
                        </Tabs>
                    </Box>
                </Grid>

                <Grid item container sx={12} spacing={10}>
                    {
                        popularCourses.map(course => (
                            <Grid item xs={12} md={6}>
                                <GradientCard
                                    description={
                                        <Fragment>
                                            <Typography variant="body" color={"primary"}>
                                                <Typography variant="body1">{course.chapters}</Typography>
                                                Chapters
                                            </Typography>
                                        </Fragment>
                                    }
                                    btnText="Explore"
                                    dpos="center"
                                >
                                    <CardContent>
                                        <Box className="CardImage">
                                            <Image
                                                src={course.image}
                                                alt={course.title}
                                                width={80}
                                                height={80}
                                                style={{
                                                    width: "100%",
                                                    height: "auto",
                                                    maxHeight: "120px",
                                                    objectFit: "cover",
                                                    borderRadius: "1rem"
                                                }}
                                            />
                                        </Box>

                                        <Typography variant="h5" sx={{ mt: 2 }}>{course.title}</Typography>
                                        <Typography variant="body" sx={{ mt: 1 }}>{course.description}</Typography>
                                    </CardContent>
                                </GradientCard>
                            </Grid>
                        ))
                    }
                </Grid>
            </Grid>

            <Grid item xs={12} md={4}>
                <Card sx={{ position: "sticky", top: "1rem" }}>
                    <CardContent>
                        <Grid container spacing={2}>
                            <Grid item xs={12}>
                                <Box
                                    sx={{
                                        position: "relative",
                                        display: "flex",
                                        alignItems: "start",
                                        gap: "1rem",
                                    }}
                                >
                                    <Box
                                        sx={{
                                            position: "relative",
                                            width: "5rem",
                                            height: "5rem",
                                            borderRadius: "1rem",
                                            overflow: "hidden",
                                            display: "flex",
                                            justifyContent: "center",
                                            alignItems: "center",
                                            backgroundColor: theme => `${theme.palette.background.default}`,
                                            boxShadow: "0 0 1rem 0 rgba(0, 0, 0, 0.1)",
                                            "& img": {
                                                maxWidth: "100%",
                                                maxHeight: "100%",
                                                width: "100%",
                                                height: "auto",
                                                objectFit: "cover",
                                            }
                                        }}
                                    >
                                        <img src="https://picsum.photos/300" alt="Astronout sitting" />
                                    </Box>

                                    <Box sx={{ width: 'calc(100% - 5rem - 1rem)', position: "relative" }}>
                                        <Tooltip title={copy ? "Copied!" : "Click to copy"} placement="top" color={copy && "success"} arrow>
                                            <Typography variant="body1"
                                                sx={{
                                                    // fazlasına üç nokta koy
                                                    width: "calc(100%)",
                                                    webkitLineClamp: 1,
                                                    display: "inline-block",
                                                    cursor: "default"
                                                }}
                                                onClick={() => {
                                                    copyToClipboard("0xRbcFu4kY0m4n0xRbcFu4kY0m4n0xRbcFu4kY0m4n")
                                                }}
                                            >
                                                0xRbcFu4kY0m4n0xRbcFu4kY0m4n0xRbcFu4kY0m4n
                                            </Typography>
                                        </Tooltip>

                                        <Typography variant="caption2" color={"secondary"}>0 / 100 XP</Typography>
                                    </Box>

                                    <LevelBar level={1} proggress={44} />
                                </Box>
                            </Grid>

                            <Grid item container xs={12} spacing={2}>
                                {
                                    statistics.map(stat => (
                                        <Grid item xs={6}>
                                            <Card variant="outlined" mode="dark" round="lg">
                                                <CardContent size="small">
                                                    <Box
                                                        sx={{
                                                            display: "flex",
                                                            alignItems: "center",
                                                        }}
                                                    >
                                                        <img src={stat.img} alt={stat.title}
                                                            style={{
                                                                width: "auto",
                                                                height: "3rem",
                                                                maxWidth: "3rem",
                                                                objectFit: "cover",
                                                                marginRight: "0.5rem"
                                                            }}
                                                        />

                                                        <Box>
                                                            <Typography variant="body1" color={"secondary"}>{stat.title}</Typography>
                                                            <Typography variant="body">{stat.value}</Typography>
                                                        </Box>
                                                    </Box>
                                                </CardContent>
                                            </Card>
                                        </Grid>
                                    ))
                                }
                            </Grid>

                            <Grid item xs={12}>
                                <Divider />
                            </Grid>

                            <Grid item container xs={12} spacing={1}>
                                {navigationData.map((item) => (
                                    <Grid item xs={12}>
                                        <ListItemButton
                                            key={item.label}
                                            sx={{ py: 0, minHeight: 32, color: 'rgba(255,255,255,.8)' }}
                                            variant="btn"
                                        >
                                            <ListItemIcon sx={{ color: 'inherit' }}>
                                                {item.icon}
                                            </ListItemIcon>
                                            <ListItemText
                                                primary={item.label}
                                                primaryTypographyProps={{ fontSize: 14, fontWeight: 'medium' }}
                                            />
                                        </ListItemButton>
                                    </Grid>
                                ))}
                            </Grid>
                        </Grid>
                    </CardContent>
                </Card>
            </Grid>
        </Grid>
    )
}

export default Courses