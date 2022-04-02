import moment from "moment"
import { shallowEqual } from "react-redux"

import { useAppSelector } from "../../redux/hooks"

export default function LicenseWarningFooter(): JSX.Element {
    const { license } = useAppSelector(state => state.main, shallowEqual)

    if (!license) {
        return <></>
    }

    let expired = false
    let message = ""
    
    if (license.kind === "trial") {
        expired = license.memberCount >= license.memberLimit
        message = "There is no more seats. You need to purchase the license."
    } else if (license.kind === "standard") {
        if (license.memberCount >= license.memberLimit) {
            expired = true
            message = "There is no more seats. You need to purchase more seats."
        } else if (moment(license.expiredAt).isBefore(new Date())) {
            expired = true
            message = "The license is expired. You need to renew the license."
        }
    }

    return (
        <>
            {(expired)? 
                <p 
                    style={{
                        textAlign: "center",
                        backgroundColor: "#fff1f0",
                        color: "#ff7875",
                        fontSize: "17px",
                        padding: "20px",
                    }}>
                    {message}
                </p>
                : 
                <></>
            }
        </>
    )
}