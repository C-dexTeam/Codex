import { coursesColumns } from '@/@local/table/columns/courses'
import ClassicTable from '@/components/tables/ClassicTable'
import { fetchCourses, getCourses } from '@/store/admin/courses'
import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'

const CoursesAdminPage = () => {

    // ** Hooks
    const dispatch = useDispatch()
    const router = useRouter()

    // ** Data
    const courses = useSelector(getCourses)

    useEffect(() => {
        dispatch(fetchCourses())
    }, [])

    return (
        <div>
            <ClassicTable
                header={{
                    title: 'Courses',
                    btnText: "Add Course",
                    btnClick: () => { router.push('/admin/courses/add') },
                }}
                // pagination={{
                //     page: filters?.page,
                //     pageCount: announcement?.pages,
                //     setPage: v => _setFilters({ ...filters, page: v }),
                // }}
                rows={courses}
                columns={coursesColumns}
            />
        </div>
    )
}


CoursesAdminPage.acl = {
    action: 'read',
    permission: 'admin'
}
CoursesAdminPage.admin = true
export default CoursesAdminPage