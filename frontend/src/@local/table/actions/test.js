import { IconButton, Tooltip } from '@mui/material'
import { useRouter } from 'next/router'
import { Delete, Edit, Visibility } from '@mui/icons-material'
import DeleteDialog from '@/components/dialog/DeleteDialog'
import { useState } from 'react'
import { deleteCourse } from '@/store/admin/courses'
import { useDispatch } from 'react-redux'
import DvrOutlinedIcon from '@mui/icons-material/DvrOutlined';
import { deleteTest } from '@/store/admin/test'

const TestActions = ({ row }) => {
    const router = useRouter()
    const [openDelete, setOpenDelete] = useState(false)

    const dispatch = useDispatch()

    console.log(row)

    const handleEdit = () => {
        router.push(`/admin/test/edit/${row.id}`)
    }


    const handleDelete = () => {
        dispatch(deleteTest(row.id))
    }

    return (
        <>
            <Tooltip title="Edit">
                <IconButton onClick={handleEdit}>
                    <Edit />
                </IconButton>
            </Tooltip>

            <Tooltip title="Delete">
                <IconButton onClick={() => setOpenDelete(true)}>
                    <Delete />
                </IconButton>
            </Tooltip>
            <DeleteDialog
                open={openDelete}
                setOpen={setOpenDelete}
                title="Delete Test"
                handleDelete={handleDelete}
            />
        </>
    )
}

export default TestActions