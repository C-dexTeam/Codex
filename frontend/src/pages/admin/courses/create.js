import CourseCreate from '@/views/courses/create'

const CourseCreatePage = () => <CourseCreate />

CourseCreatePage.acl = {
    action: 'read',
    permission: 'admin'
}
CourseCreatePage.admin = true
export default CourseCreatePage 