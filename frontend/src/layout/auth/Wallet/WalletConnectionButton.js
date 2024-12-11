import { BaseWalletMultiButton } from '@solana/wallet-adapter-react-ui'

const WalletConnectionButton = () => {
    return <BaseWalletMultiButton
        labels={{
            'change-wallet': 'Change wallet',
            connecting: 'Connecting ...',
            'copy-address': 'Copy address',
            copied: 'Copied',
            disconnect: 'Disconnect',
            'has-wallet': 'Connect',
            'no-wallet': 'Connect Wallet',
        }}
    />
}

export default WalletConnectionButton