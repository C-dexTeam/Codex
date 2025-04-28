import TetsCreate from "@/views/admin/test/create"


const TetsCreatePage = () => <TetsCreate />

TetsCreatePage.acl = {
    action: 'read',
    permission: 'admin'
}
TetsCreatePage.admin = true
export default TetsCreatePage