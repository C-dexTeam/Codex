import { courseValues } from "@/@local/table/form-values/event/defaultValues"
import AddCourseForm from "@/components/form/course/add"
import { createCourse } from "@/store/admin/courses"
import { Card, CardContent, Grid, Typography } from "@mui/material"
import { useRouter } from "next/router"
import { useState } from "react"
import { useDispatch } from "react-redux"

const CourseAdd = () => {
    const [values, setValues] = useState(courseValues)

    const dispatch = useDispatch()
    const router = useRouter()

    const handleSubmit = (formData) => {
        dispatch(createCourse({ formData, callback: () => router.replace("/admin/courses") }))

    }

    return (
        <Grid container spacing={2}>
            <Grid item xs={12} md={12}>
                <Typography variant="h2">Create Course</Typography>
            </Grid>

            <Grid item xs={12} md={8}>
                <Card>
                    <CardContent>
                        <AddCourseForm
                            values={values}
                            setValues={setValues}
                            handleSubmit={handleSubmit}
                        />
                    </CardContent>
                </Card>
            </Grid>

            <Grid item xs={12} md={4}>
                {JSON.stringify(values)}
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