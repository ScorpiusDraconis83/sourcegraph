import React, { useMemo, type FC } from 'react'

import { parseISO } from 'date-fns'
import type { Observable } from 'rxjs'
import { catchError, map } from 'rxjs/operators'

import { asError, type ErrorLike, isErrorLike, numberWithCommas } from '@sourcegraph/common'
import { gql, dataOrThrowErrors } from '@sourcegraph/http-client'
import {
    LoadingSpinner,
    useObservable,
    Link,
    CardFooter,
    Alert,
    ButtonLink,
    Tooltip,
    ErrorAlert,
    Text,
} from '@sourcegraph/wildcard'

import { queryGraphQL } from '../../../backend/graphql'
import type { ProductLicenseInfoResult } from '../../../graphql-operations'
import { formatUserCount } from '../../../productSubscription/helpers'
import { ExpirationDate } from '../../productSubscription/ExpirationDate'
import { ProductCertificate } from '../../productSubscription/ProductCertificate'
import { TrueUpStatusSummary } from '../../productSubscription/TrueUpStatusSummary'

const queryProductLicenseInfo = (): Observable<{
    productSubscription: ProductLicenseInfoResult['site']['productSubscription']
    currentUserCount: number
}> =>
    queryGraphQL<ProductLicenseInfoResult>(gql`
        query ProductLicenseInfo {
            site {
                productSubscription {
                    productNameWithBrand
                    actualUserCount
                    actualUserCountDate
                    noLicenseWarningUserCount
                    license {
                        ...ProductLicenseInfoLicenseFields
                    }
                }
            }
            users {
                totalCount
            }
        }
        fragment ProductLicenseInfoLicenseFields on ProductLicenseInfo {
            tags
            userCount
            expiresAt
            isValid
            licenseInvalidityReason
        }
    `).pipe(
        map(dataOrThrowErrors),
        map(({ site, users }) => ({
            productSubscription: site.productSubscription,
            currentUserCount: users.totalCount,
        }))
    )

interface Props {
    className?: string

    /**
     * If true, always show the license true-up status.
     * If undefined or false, never show the full license true-up status, and instead only show an alert
     * if the user count is over the license limit.
     *
     */
    showTrueUpStatus?: boolean
}

/**
 * A component displaying information about and the status of the product subscription.
 */
export const ProductSubscriptionStatus: React.FunctionComponent<React.PropsWithChildren<Props>> = ({
    className,
    showTrueUpStatus,
}) => {
    /** The product subscription status, or an error, or undefined while loading. */
    const statusOrError = useObservable(
        useMemo(() => queryProductLicenseInfo().pipe(catchError((error): [ErrorLike] => [asError(error)])), [])
    )
    if (statusOrError === undefined) {
        return (
            <div className="text-center">
                <LoadingSpinner />
            </div>
        )
    }
    if (isErrorLike(statusOrError)) {
        return <ErrorAlert error={statusOrError} prefix="Error checking product license" />
    }

    const {
        productSubscription: {
            productNameWithBrand,
            actualUserCount,
            actualUserCountDate,
            noLicenseWarningUserCount,
            license,
        },
        currentUserCount,
    } = statusOrError

    // No license means Sourcegraph Free. For that, show the user that they can use this for free
    // forever, and show them how to upgrade.

    return (
        <div>
            <ProductCertificate
                title={productNameWithBrand}
                detail={<LicenseDetails license={license} />}
                footer={
                    <CardFooter className="d-flex align-items-center justify-content-between">
                        {license?.isValid ? (
                            <>
                                <div>
                                    <strong>User licenses:</strong> {numberWithCommas(currentUserCount)} currently used
                                    / {numberWithCommas(license.userCount - currentUserCount)} remaining (
                                    {numberWithCommas(actualUserCount)} maximum ever used)
                                </div>
                                <ButtonLink
                                    to="https://sourcegraph.com/pricing"
                                    target="_blank"
                                    rel="noopener"
                                    variant="primary"
                                    size="sm"
                                >
                                    Upgrade
                                </ButtonLink>
                            </>
                        ) : (
                            <>
                                <div className="mr-2">
                                    Add a license key to activate Sourcegraph Enterprise features{' '}
                                    {typeof noLicenseWarningUserCount === 'number'
                                        ? `or to exceed ${noLicenseWarningUserCount} users`
                                        : ''}
                                </div>
                                <div className="text-nowrap flex-wrap-reverse">
                                    <Tooltip content="Buy a Sourcegraph Enterprise subscription to get a license key">
                                        <ButtonLink
                                            to="http://sourcegraph.com/contact/sales"
                                            target="_blank"
                                            rel="noopener"
                                            variant="primary"
                                            size="sm"
                                        >
                                            Get license
                                        </ButtonLink>
                                    </Tooltip>
                                </div>
                            </>
                        )}
                    </CardFooter>
                }
                className={className}
            />
            {license &&
                (showTrueUpStatus ? (
                    <TrueUpStatusSummary
                        actualUserCount={actualUserCount}
                        actualUserCountDate={actualUserCountDate}
                        license={license}
                    />
                ) : (
                    license.userCount - actualUserCount < 0 && (
                        <Alert variant="warning">
                            You have exceeded your licensed users.{' '}
                            <Link to="/site-admin/license">View your license details</Link> or{' '}
                            <Link to="https://sourcegraph.com/pricing" target="_blank" rel="noopener">
                                upgrade your license
                            </Link>{' '}
                            to true up and prevent a retroactive charge.
                        </Alert>
                    )
                ))}
        </div>
    )
}

interface LicenseDetailsProps {
    license: ProductLicenseInfoResult['site']['productSubscription']['license']
}

const LicenseDetails: FC<LicenseDetailsProps> = ({ license }) => {
    if (!license) {
        return null
    }

    if (license.isValid) {
        return (
            <>
                {formatUserCount(license.userCount, true)} license,{' '}
                <ExpirationDate
                    date={parseISO(license.expiresAt)}
                    showRelative={true}
                    lowercase={true}
                    showPrefix={true}
                />
            </>
        )
    }

    return (
        <Alert variant="danger">
            <Text className="mb-0">
                The Sourcegraph license key is invalid. Reason: {license.licenseInvalidityReason}
            </Text>
        </Alert>
    )
}
