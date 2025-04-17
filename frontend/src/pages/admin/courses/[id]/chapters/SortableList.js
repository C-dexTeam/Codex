import React from 'react';
import { DndContext, closestCenter } from '@dnd-kit/core';
import { SortableContext, verticalListSortingStrategy } from '@dnd-kit/sortable';
import { DraggableChapter } from './DraggableChapter';
import {
    restrictToVerticalAxis,
} from '@dnd-kit/modifiers';

export function SortableList({ chapters = [], onSortEnd }) {

    return (
        <DndContext
            collisionDetection={closestCenter}
            modifiers={[restrictToVerticalAxis]}
            onDragEnd={onSortEnd}
        >
            <SortableContext
                items={chapters.map(chapter => chapter.id)}
                strategy={verticalListSortingStrategy}
            >
                {chapters.map(chapter => (
                    <DraggableChapter
                        key={chapter.id}
                        id={chapter.id}
                        chapter={chapter}
                    />
                ))}
            </SortableContext>
        </DndContext>
    );
}