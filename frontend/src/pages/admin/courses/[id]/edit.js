import React from 'react'
import CourseEdit from './index'

const CourseEditPage = () => <CourseEdit />

CourseEditPage.acl = {
    action: 'read',
    permission: 'admin'
}
CourseEditPage.admin = true
export default CourseEditPage