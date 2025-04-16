import { useState, useEffect } from 'react'
import { useRouter } from 'next/router'
import { useDispatch, useSelector } from 'react-redux'
import { Box, Button, Card, CardContent, Typography, Grid, IconButton, Popover } from '@mui/material'
import { DragIndicator, FilterList, Add } from '@mui/icons-material'
import { fetchChapters, getChapters } from '@/store/admin/chapters'
import { DndProvider } from 'react-dnd'
import { HTML5Backend } from 'react-dnd-html5-backend'
import { useDrag, useDrop } from 'react-dnd'
import ChapterFilter from '@/components/filter/ChapterFilter'
import { showToast } from '@/utils/showToast'

const ChapterCard = ({ chapter, index, moveChapter }) => {
    const [{ isDragging }, drag] = useDrag({
        type: 'CHAPTER',
        item: { id: chapter.id, index },
        collect: (monitor) => ({
            isDragging: monitor.isDragging(),
        }),
    })

    const [, drop] = useDrop({
        accept: 'CHAPTER',
        hover: (draggedItem) => {
            if (draggedItem.index !== index) {
                moveChapter(draggedItem.index, index)
                draggedItem.index = index
            }
        },
    })

    return (
        <Card
            ref={(node) => drag(drop(node))}
            sx={{
                mb: 2,
                opacity: isDragging ? 0.5 : 1,
                cursor: 'move',
                display: 'flex',
                alignItems: 'center',
                p: 2
            }}
        >
            <DragIndicator sx={{ mr: 2, color: 'text.secondary' }} />
            <CardContent sx={{ flex: 1 }}>
                <Typography variant="h6">{chapter.title}</Typography>
                <Typography variant="body2" color="text.secondary">
                    {chapter.description}
                </Typography>
            </CardContent>
        </Card>
    )
}

const CourseChapters = () => {
    const router = useRouter()
    const dispatch = useDispatch()
    const chapters = useSelector(getChapters)
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

    useEffect(() => {
        if (router.query.id) {
            dispatch(fetchChapters({ params: filters }))
        }
    }, [router.query.id, filters])

    const handleFilterClick = (event) => {
        setFilterAnchor(event.currentTarget)
    }

    const handleFilterClose = () => {
        setFilterAnchor(null)
    }

    const moveChapter = (fromIndex, toIndex) => {
        // TODO: Implement chapter reordering API call
        showToast('info', 'Chapter order updated')
    }

    return (
        <Box sx={{ p: 3 }}>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
                <Typography variant="h4">Course Chapters</Typography>
                <Box>
                    <IconButton onClick={handleFilterClick} sx={{ mr: 2 }}>
                        <FilterList />
                    </IconButton>

                    <Button
                        variant="contained"
                        startIcon={<Add />}
                        onClick={() => router.push(`/admin/courses/${router.query.id}/chapters/create`)}
                    >
                        Create Chapter
                    </Button>
                </Box>
            </Box>

            <DndProvider backend={HTML5Backend}>
                <Grid container spacing={3}>
                    <Grid item xs={12}>
                        {chapters?.data?.map((chapter, index) => (
                            <ChapterCard
                                key={chapter.id}
                                chapter={chapter}
                                index={index}
                                moveChapter={moveChapter}
                            />
                        ))}
                    </Grid>
                </Grid>
            </DndProvider>

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