import { coursesColumns } from '@/@local/table/columns/courses'
import ClassicTable from '@/components/tables/ClassicTable'

const CoursesAdminPage = () => {
    return (
        <div>
            <ClassicTable
                header={{
                    title: 'Courses',
                    btnText: "Add Course",
                    btnClick: () => { router.push('/courses/add') },
                }}
                // pagination={{
                //     page: filters?.page,
                //     pageCount: announcement?.pages,
                //     setPage: v => _setFilters({ ...filters, page: v }),
                // }}
                rows={[
                    {
                        id: 1,
                        email: "user1@example.com",
                        createdDate: new Date('2023-01-01').toISOString()
                    },
                    {
                        id: 2,
                        email: "user2@example.com",
                        createdDate: new Date('2023-02-01').toISOString()
                    },
                    {
                        id: 3,
                        email: "user3@example.com",
                        createdDate: new Date('2023-03-01').toISOString()
                    },
                    {
                        id: 4,
                        email: "user4@example.com",
                        createdDate: new Date('2023-04-01').toISOString()
                    }
                ]}
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