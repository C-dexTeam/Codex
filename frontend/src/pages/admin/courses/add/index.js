import AddCourseForm from "@/components/form/course/add"
import { Card, CardContent, Grid, Typography } from "@mui/material"

const CourseAdd = () => {
    return (
        <Grid container spacing={2}>
            <Grid item container xs={12} md={8} spacing={2}>
                <Grid item xs={12}>
                    <Typography variant="h2">Add Course</Typography>
                </Grid>

                <Grid item xs={12}>
                    <Card>
                        <CardContent>
                            <AddCourseForm />
                        </CardContent>
                    </Card>
                </Grid>
            </Grid>

            <Grid item xs={12} md={4}>
            </Grid>
        </Grid>
    )
}

CourseAdd.acl = {
    action: 'read',
    permission: 'admin'
}
CourseAdd.admin = true
export default CourseAdd