// ** MUI Imports
import {
    Typography,
    Button,
    Grid,
    Box,
    useMediaQuery
} from '@mui/material'
import DefaultTextField from '../form/components/DefaultTextField'

const TableHeader = (props) => {
    const {
        title = null,
        searchText = null,
        handleSearch = undefined,
        btnText = null,
        btnClick = undefined,
        totalCount = null,
        totalMessage = null,
        child = null
    } = props

    const _sx = useMediaQuery(theme => theme.breakpoints.down('sm'))

    return (
        <Grid sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', py: '20px' }} container spacing={2.5}>
            <Grid item xs={12} md sm>
                {
                    title ? <Typography variant='h5' sx={{ fontWeight: '600' }}>  {title}</Typography> : null
                }
            </Grid>

            <Grid item xs={12} md={10} nm={8} sm={10}>
                <Box sx={{ display: 'flex', justifyContent: 'flex-end', alignItems: 'center', gap: '20px', flexDirection: _sx ? 'column' : undefined }}>
                    {
                        child != null
                            ? child
                            : null
                    }

                    {
                        searchText != null
                            ? <DefaultTextField
                                type='text'
                                name="search"
                                value={searchText}
                                onChange={e => handleSearch(e.target.value)}
                                placeholder="Ne arıyorsun?"
                                size='small'
                                noMargin
                            />
                            : null
                    }

                    {
                        btnText
                            ? <Button
                                fullWidth={_sx}
                                variant='outlined'
                                color='primary'
                                size='small'
                                sx={{ textTransform: 'none' }}
                                onClick={() => { btnClick() }}
                            >
                                {btnText}
                            </Button>
                            : null
                    }

                    {
                        totalCount
                            ? <Typography variant='caption' sx={{ width: 'fit-content', minWidth: `calc(16px + ${totalMessage?.length * 8}px)` }}>{totalCount} {totalMessage}</Typography>
                            : null
                    }
                </Box>
            </Grid>
        </Grid>
    )
}

export default TableHeader