
import CourseEdit from "@/views/courses/edit"

const CourseEditPage = () => <CourseEdit />

CourseEditPage.acl = {
    action: 'read',
    permission: 'admin'
}
CourseEditPage.admin = true
export default CourseEditPage 