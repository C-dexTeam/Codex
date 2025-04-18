import RewardEdit from '@/views/rewards/edit'

const RewardEditPage = () => <RewardEdit />

RewardEditPage.acl = {
    action: 'read',
    permission: 'admin'
}
RewardEditPage.admin = true
export default RewardEditPage 