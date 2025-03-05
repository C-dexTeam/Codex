import { Delete, Edit } from '@mui/icons-material'
import { IconButton } from '@mui/material'
import { useRouter } from 'next/router'
import { useState } from 'react'
import { useDispatch } from 'react-redux'

const CourseActions = ({ row }) => {
    const [openDelete, setOpenDelete] = useState(false)

    const dispatch = useDispatch()
    const router = useRouter()

    const handleRotateEdit = () => {
        router.push(`/makale/duzenle/${row.id}`)
    }

    return (
        <>
            <IconButton size='small' onClick={() => handleRotateEdit()} color='warning'>
                <Edit />
            </IconButton>

            <IconButton size='small' onClick={() => setOpenDelete(true)} color='error'>
                <Delete />
            </IconButton>

            {/* <DeleteDialog
                open={openDelete}
                setOpen={setOpenDelete}
                title="Makaleyi silmek istediğinize emin misiniz?"
                handleDelete={handleDelete}
            /> */}
        </>
    )
}

export default CourseActions