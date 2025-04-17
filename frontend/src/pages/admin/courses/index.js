import { coursesColumns } from '@/@local/table/columns/courses'
import ClassicTable from '@/components/tables/ClassicTable'
import { fetchCourses, getCourses } from '@/store/admin/courses'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'

const CoursesAdminPage = () => {

    // ** States
    const [filters, setFilters] = useState({
        page: 1,
        limit: 10,
        title: "",
    })

    // ** Hooks
    const dispatch = useDispatch()
    const router = useRouter()

    // ** Data
    const courses = useSelector(getCourses)
    console.log(courses);
    

    useEffect(() => {
        dispatch(fetchCourses(filters))
    }, [filters])

    return (
        <div>
            <ClassicTable
                header={{
                    title: 'Courses',
                    btnText: "Create Course",
                    btnClick: () => { router.push('/admin/courses/create') },
                    search: filters?.title,
                    handleSearch: v => setFilters({ ...filters, title: v }),
                    totalCount: courses?.totalCount,
                }}
                rows={courses}
                columns={coursesColumns}
                getRowId={(row) => row._id || row.id || Math.random().toString(36).substr(2, 9)}
                pagination={{
                    page: filters?.page,
                    pageCount: Math.ceil(courses?.totalCount / filters?.limit),
                    setPage: v => setFilters({ ...filters, page: v }),
                }}
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