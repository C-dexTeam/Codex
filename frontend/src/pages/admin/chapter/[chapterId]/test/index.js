import TetsList from "@/views/admin/test/list"


const TetsListPage = () => <TetsList />

TetsListPage.acl = {
    action: 'read',
    permission: 'admin'
}
TetsListPage.admin = true
export default TetsListPage