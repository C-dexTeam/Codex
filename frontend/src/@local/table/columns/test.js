// ** MUI Imports
import { showDatetime } from "@/utils/timeOptions"
import {
    Typography
} from "@mui/material"
import CustomTooltip from "@/components/tooltip"
import TestActions from "../actions/test"

export const testColums = [
    {
        flex: 0.1,
        minWidth: 152,
        headerName: "",
        field: "actions",
        renderCell: params => <TestActions row={params.row} />
    },
    {
        flex: 0.55,
        minWidth: 100,
        headerName: 'Input Value',
        field: 'input',
        renderCell: params => {
            const { row } = params

            return (
                <CustomTooltip title={row.input} placement='top'>
                    <Typography variant='body1' sx={{ cursor: 'default', overflow: 'hidden', textOverflow: 'ellipsis' }}>
                        {row.input}
                    </Typography>
                </CustomTooltip>
            )
        }
    },
    {
        flex: 0.55,
        minWidth: 100,
        headerName: 'Output Value',
        field: 'output',
        renderCell: params => {
            const { row } = params

            return (
                <CustomTooltip title={row.output} placement='top'>
                    <Typography variant='body1' sx={{ cursor: 'default', overflow: 'hidden', textOverflow: 'ellipsis' }}>
                        {row.output}
                    </Typography>
                </CustomTooltip>
            )
        }
    },
]