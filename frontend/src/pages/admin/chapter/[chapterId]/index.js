import ChapterEdit from '@/views/chapters/edit'

const ChapterEditPage = () => <ChapterEdit />

ChapterEditPage.acl = {
    action: 'read',
    permission: 'admin'
}
ChapterEditPage.admin = true
export default ChapterEditPage