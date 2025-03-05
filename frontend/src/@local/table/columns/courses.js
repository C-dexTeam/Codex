// ** MUI Imports
import { showDatetime } from "@/utils/timeOptions"
import {
    Tooltip,
    Typography
} from "@mui/material"
import CourseActions from "../actions/article"

export const coursesColumns = [
    // {
    //     flex: 0.03,
    //     minWidth: 30,
    //     headerName: '',
    //     field: 'actions',
    //     renderCell: params => <SubscriptionsActions row={params?.row} />
    // },
    {
        flex: 0.02,
        minWidth: 20,
        headerName: "",
        field: "actions",
        renderCell: params => <CourseActions row={params.row} />
    },
    {
        flex: 0.1,
        minWidth: 100,
        headerName: 'E-posta',
        field: 'email',
        renderCell: params => {
            const { row } = params

            return (
                <Tooltip title={row.fullname} placement='top'>
                    <Typography variant='body1' sx={{ cursor: 'default', overflow: 'hidden', textOverflow: 'ellipsis' }}>
                        {row.email}
                    </Typography>
                </Tooltip>
            )
        }
    },
    {
        flex: 0.1,
        minWidth: 100,
        headerName: 'ABONELİK TARİHİ',
        field: 'createdDate',
        renderCell: params => {
            const { row } = params

            return (
                <Typography variant='body1' sx={{ cursor: 'default' }}>
                    {showDatetime({ date: row.createdDate }) ?? '-'}
                </Typography>
            )
        }
    },
]