
import { rewardsColumns } from '@/@local/table/columns/rewards'
import ClassicTable from '@/components/tables/ClassicTable'
import { fetchRewards, getRewards } from '@/store/admin/rewards'
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
    const rewards = useSelector(getRewards)

    // ** Effects
    useEffect(() => {
        dispatch(fetchRewards(filters))
    }, [filters])

    return (
        <div>
            <ClassicTable
                header={{
                    title: 'Rewards',
                    btnText: "Create Reward",
                    btnClick: () => { router.push('/admin/rewards/create') },
                    search: filters?.title,
                    handleSearch: v => setFilters({ ...filters, title: v }),
                    totalCount: rewards?.totalCount,
                }}
                rows={rewards || []}
                columns={rewardsColumns}
                getRowId={(row) => row._id || row.id || Math.random().toString(36).substr(2, 9)}
                pagination={{
                    page: filters?.page,
                    pageCount: Math.ceil(rewards?.totalCount / filters?.limit),
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