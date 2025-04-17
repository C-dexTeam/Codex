import React from 'react';
import { useSortable } from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';
import { Card, CardContent, Typography, Button, Box, CardActions } from '@mui/material';
import { useRouter } from 'next/router';
import { DragIndicator } from '@mui/icons-material';

export function DraggableChapter({ id, chapter }) {
    const router = useRouter();
    const {
        attributes,
        listeners,
        setNodeRef,
        transform,
        transition,
        isDragging
    } = useSortable({ id });

    const style = {
        transform: CSS.Transform.toString(transform),
        transition,
        opacity: isDragging ? 0.5 : 1,
        cursor: "move",
        display: "flex",
    };

    const handleDetailsClick = () => {
        router.push(`/admin/courses/${chapter.course_id}/chapters/${chapter.id}`);
    };

    const DragHandle = () => (
        <Box
            style={{
                width: '100%',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                cursor: 'grab',
            }}
            onMouseEnter={(e) => { if (e.target.style.cursor !== 'grab') e.target.style.cursor = 'grabbing' }}
            onMouseLeave={(e) => { if (e.target.style.cursor === 'grabbing') e.target.style.cursor = '' }}
            onMouseDown={(e) => { e.target.style.cursor = 'grabbing' }}
        >
            <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                <DragIndicator sx={{ mb: -0.75, color: theme => theme.palette.border.light }} />
                <DragIndicator sx={{ mt: -0.75, color: theme => theme.palette.border.light }} />
                <DragIndicator sx={{ mt: -0.75, color: theme => theme.palette.border.light }} />
            </Box>
        </Box>
    );

    return (
        <Card
            ref={setNodeRef}
            style={style}
            {...attributes}
            {...listeners}
            sx={{ mb: 2 }}
        >
            <Box
                sx={{
                    display: 'flex',
                    alignItems: 'center',
                    borderRight: theme => `1px solid ${theme.palette.border.secondary}`,
                    '&:hover': {
                        backgroundColor: theme => theme.palette.action.hover
                    }
                }}
            >
                <DragHandle />
            </Box>

            <CardContent sx={{ flex: 1, padding: "1rem !important" }}>
                <Typography variant="subtitle">{chapter.title}</Typography>
                <Typography variant="body">{chapter.description}</Typography>
            </CardContent>

            <CardActions flat>
                <Button
                    onClick={handleDetailsClick}
                    variant='text'
                    color='secondary'
                >
                    View
                </Button>
            </CardActions>
        </Card>
    );
}