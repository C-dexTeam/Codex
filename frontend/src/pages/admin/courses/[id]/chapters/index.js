import { useState, useEffect } from 'react'
import { useRouter } from 'next/router'
import { useDispatch, useSelector } from 'react-redux'
import { Box, Button, Typography, IconButton, Popover } from '@mui/material'
import { FilterList, Add } from '@mui/icons-material'
import { fetchChapters, getChapters, updateChapter } from '@/store/admin/chapters'
import ChapterFilter from '@/components/filter/ChapterFilter'
import { arrayMove } from '@dnd-kit/sortable'
import { SortableList } from './SortableList'

const CourseChapters = () => {
    const router = useRouter()
    const dispatch = useDispatch()

    const [filterAnchor, setFilterAnchor] = useState(null)
    const [filters, setFilters] = useState({
        page: 1,
        limit: 10,
        courseID: router.query.id,
        title: '',
        languageID: '',
        rewardID: '',
        grantsExperience: '',
        active: ''
    })

    const chapters = useSelector(getChapters)

    const handleFilterClick = (event) => {
        setFilterAnchor(event.currentTarget)
    }

    const handleFilterClose = () => {
        setFilterAnchor(null)
    }

    const handleSortEnd = (event) => {
        const { active, over } = event;
        console.log("qweqwe", active, over);

        if (active?.data?.current?.sortable?.index !== over?.data?.current?.sortable?.index) {
            const oldIndex = chapters.findIndex((item, index) => index === active?.data?.current?.sortable?.index);
            const newIndex = chapters.findIndex((item, index) => index === over?.data?.current?.sortable?.index);
            const newItems = arrayMove(chapters, oldIndex, newIndex);

            newItems?.forEach((item, index) => {
                dispatch(updateChapter({ ...item, order: index }))
            });
        }
    }

    useEffect(() => {
        if (router.query.id) {
            dispatch(fetchChapters({ params: filters }))
        }
    }, [router.query.id, filters])

    return (
        <Box>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
                <Typography variant="h4">Course Chapters</Typography>
                <Box>
                    <IconButton onClick={handleFilterClick} sx={{ mr: 2 }}>
                        <FilterList />
                    </IconButton>

                    <Button
                        variant="outlined"
                        color="secondary"
                        startIcon={<Add />}
                        onClick={() => router.push(`/admin/courses/${router.query.id}/chapters/create`)}
                    >
                        Create Chapter
                    </Button>
                </Box>
            </Box>

            <SortableList
                chapters={chapters}
                onSortEnd={handleSortEnd}
            />
            {/* <Grid container spacing={3}>
                {chapters?.data?.map((chapter, index) => (
                    <Grid item xs={12} key={chapter.id}>
                        <ChapterCard chapter={chapter} index={index} />
                    </Grid>
                ))}
            </Grid> */}

            <Popover
                open={Boolean(filterAnchor)}
                anchorEl={filterAnchor}
                onClose={handleFilterClose}
                anchorOrigin={{
                    vertical: 'bottom',
                    horizontal: 'right',
                }}
                transformOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                }}
            >
                <ChapterFilter
                    filters={filters}
                    setFilters={setFilters}
                    onClose={handleFilterClose}
                />
            </Popover>
        </Box>
    )
}

CourseChapters.acl = {
    action: 'read',
    permission: 'admin'
}
CourseChapters.admin = true
export default CourseChapters
