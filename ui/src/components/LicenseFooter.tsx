import moment from "moment"
import { License } from "../models"

interface LicenseFooterProps {
    license?: License
}

export default function LicenseFooter(props: LicenseFooterProps): JSX.Element {
    if (!props.license) {
        return <div></div>
    }

    let expired = false
    let message = ""
    
    if (props.license.kind === "trial") {
        expired = props.license.memberCount >= props.license.memberLimit
        message = "There is no more seats. You need to purchase the license."
    } else if (props.license.kind === "standard") {
        if (props.license.memberCount >= props.license.memberLimit) {
            expired = true
            message = "There is no more seats. You need to purchase more seats."
        } else if (moment(props.license.expiredAt).isBefore(new Date())) {
            expired = true
            message = "The license is expired. You need to renew the license."
        }
    }

    return (
        <div>
            {(expired)? 
                <p 
                    style={{
                        textAlign: "center",
                        backgroundColor: "#fff1f0",
                        color: "#ff7875",
                        fontSize: "17px",
                        padding: "20px",
                    }}
                >{message}</p>: 
                null
            }
        </div>
    )
}