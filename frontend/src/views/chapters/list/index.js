import { useState, useEffect } from 'react'
import { useRouter } from 'next/router'
import { useDispatch, useSelector } from 'react-redux'
import { Box, Button, Typography, IconButton, Popover } from '@mui/material'
import { Add } from '@mui/icons-material'
import { fetchChapters, getChapters, updateChapter } from '@/store/admin/chapters'
import { arrayMove } from '@dnd-kit/sortable'
import { fetchCourse, getCourse } from '@/store/admin/courses'
import DefaultTextField from '@/components/form/components/DefaultTextField'
import SortableList from '@/components/dnd/SortableList'
import CustomBreadcrumbs from '@/components/breadcrumbs'
const CourseChaptersList = () => {
    // ** Hooks
    const router = useRouter()
    const dispatch = useDispatch()

    // ** States
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

    // ** Selectors
    const chapters = useSelector(getChapters)
    const course = useSelector(getCourse)

    const handleSortEnd = (event) => {
        const { active, over } = event;

        if (active?.data?.current?.sortable?.index !== over?.data?.current?.sortable?.index) {
            const oldIndex = chapters.findIndex((item, index) => index === active?.data?.current?.sortable?.index);
            const newIndex = chapters.findIndex((item, index) => index === over?.data?.current?.sortable?.index);
            const newItems = arrayMove(chapters, oldIndex, newIndex);

            newItems?.forEach((item, index) => {
                dispatch(updateChapter({ ...item, order: index }))
            });
        }
    }

    // ** Effects
    useEffect(() => {
        if (router.query.id) {
            dispatch(fetchChapters({ params: filters }))
            dispatch(fetchCourse(router.query.id))
        }
    }, [router.query.id, filters])

    return (
        <Box>
            <CustomBreadcrumbs
                titles={[
                    { title: 'Admin', path: '/admin' },
                    { title: 'Courses', path: '/admin/courses' },
                    { title: 'Create Course' }
                ]}
            />

            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
                <Typography variant="h4">
                    {course?.title} -
                    <Typography variant="h6">
                        {chapters?.totalCount} chapters
                    </Typography>
                </Typography>

                <Box sx={{ display: 'flex', gap: 2 }}>
                    <DefaultTextField
                        value={filters.title}
                        onChange={(e) => setFilters({ ...filters, title: e.target.value })}
                        noMargin
                        noControl
                        sx={{
                            height: "100%",
                            "& .MuiInputBase-root": {
                                height: "100%",
                            },
                        }}
                        placeholder="Search chapter"
                    />

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

            {/* <Popover
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
            </Popover> */}
        </Box>
    )
}

export default CourseChaptersList