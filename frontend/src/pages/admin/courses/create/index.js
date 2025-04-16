import { courseValues } from "@/@local/table/form-values/courses/defaultValues"
import CourseForm from "@/components/form/course/form"
import { createCourse } from "@/store/admin/courses"
import { Card, CardContent, Grid, Typography } from "@mui/material"
import { useRouter } from "next/router"
import { useState } from "react"
import { useDispatch } from "react-redux"
import CustomBreadcrumbs from "@/components/breadcrumbs"
import CourseCardPreview from "@/components/card/CourseCardPreview"

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
                <CustomBreadcrumbs
                    titles={[
                        { title: 'Admin', path: '/admin' },
                        { title: 'Courses', path: '/admin/courses' },
                        { title: 'Create Course' }
                    ]}
                />
                <Typography variant="h2" sx={{ mt: 2 }}>Create Course</Typography>
            </Grid>

            <Grid item xs={12} md={8}>
                <Card>
                    <CardContent>
                        <CourseForm
                            values={values}
                            setValues={setValues}
                            handleSubmit={handleSubmit}
                        />
                    </CardContent>
                </Card>
            </Grid>

            <Grid item xs={12} md={4}>
                <CourseCardPreview values={values} />
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