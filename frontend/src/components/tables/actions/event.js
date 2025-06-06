import { IconButton } from '@mui/material'
import { PencilOutline, TrashCanOutline } from 'mdi-material-ui'
import { useRouter } from 'next/router'
import { useState } from 'react'
import { useDispatch } from 'react-redux'
import DeleteDialog from 'src/components/dialogs/DeleteDialog'
import { deleteContent } from 'src/store/api/content'

const EventActions = ({ row }) => {
    const [openDelete, setOpenDelete] = useState(false)

    const dispatch = useDispatch()
    const router = useRouter()

    const handleDelete = () => {
        dispatch(deleteContent({ id: row.id, type: "event" }))
            .finally(() => setOpenDelete(false))
    }

    const handleRotateEdit = () => {
        router.push(`/etkinlik/duzenle/${row.id}`)
    }

    return (
        <>
            <IconButton size='small' onClick={() => handleRotateEdit()} color='warning'>
                <PencilOutline />
            </IconButton>

            <IconButton size='small' onClick={() => setOpenDelete(true)} color='error'>
                <TrashCanOutline />
            </IconButton>

            <DeleteDialog
                open={openDelete}
                setOpen={setOpenDelete}
                title="Etkinliği silmek istediğinize emin misiniz?"
                handleDelete={handleDelete}
            />
        </>
    )
}

export default EventActions