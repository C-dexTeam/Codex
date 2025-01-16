import GradientCard from "@/components/card/GradientCard"
import { Box, Card, CardContent, Grid, Typography } from "@mui/material"
import Image from "next/image"
import { Fragment } from "react"

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

const Courses = () => {
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
                            top: 0,
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
            </Grid>

            <Grid item xs={12} md={4}>
                <Card>
                    <CardContent>
                        <Grid container spacing={2}>
                            <Grid item xs={12}>pp</Grid>
                            <Grid item xs={12}>statistics</Grid>
                            <Grid item xs={12}>divider</Grid>
                            <Grid item xs={12}>nav</Grid>
                        </Grid>
                    </CardContent>
                </Card>
            </Grid>
        </Grid>
    )
}

export default Courses