WITH cp AS (
    SELECT
        school_code,
        amount
    FROM school_reports
    WHERE scope = 'CENTRAL'
      AND section = 'PRIMARY'
      AND year = ?
      AND month = ?
),
	sp AS (
	    SELECT
        school_code,
        amount
    FROM school_reports
    WHERE scope = 'STATE'
      AND section = 'PRIMARY'
      AND year = ?
      AND month = ?
),
	cup AS (
		SELECT
		school_code,
		amount
    FROM school_reports
    WHERE scope = 'CENTRAL'
      AND section = 'UPPER_PRIMARY'
      AND year = ?
      AND month = ?
),
	sup AS (
    SELECT
		school_code,
		amount
    FROM school_reports
    WHERE scope = 'STATE'
      AND section = 'UPPER_PRIMARY'
      AND year = ?
      AND month = ?
)
SELECT
    b.block_code,
	b.block_name,
	c.cluster_code,
	c.cluster_name,
    s.school_code,
    sb.school_name,
    sb.pan_num,
    sb.tan_num,
    sb.address,
    sb.city,
    sb.country,
    sb.state,
    sb.district,
    sb.pincode,
    sb.mobile,
    sb.email,
    sb.bank_name,
    sb.ifsc_code,
    sb.account_number,

    COALESCE(cp.amount, 0) AS Central_Primary,
    COALESCE(sp.amount, 0) AS State_Primary,
    COALESCE(cp.amount, 0) + COALESCE(sp.amount, 0) AS Primary_Total,

    COALESCE(cup.amount, 0) AS Central_Upper_Primary,
    COALESCE(sup.amount, 0) AS State_Upper_Primary,
    COALESCE(cup.amount, 0) + COALESCE(sup.amount, 0) AS Upper_Primary_Total,

	COALESCE(cp.amount, 0) + COALESCE(sp.amount, 0)+COALESCE(cup.amount, 0) + COALESCE(sup.amount, 0) AS Gross_Toal,

	COALESCE((COALESCE(cp.amount, 0) + COALESCE(sp.amount, 0)+COALESCE(cup.amount, 0) + COALESCE(sup.amount, 0))*0.7340,0) AS General,
	COALESCE((COALESCE(cp.amount, 0) + COALESCE(sp.amount, 0)+COALESCE(cup.amount, 0) + COALESCE(sup.amount, 0))*0.1384,0) AS SC,
	COALESCE((COALESCE(cp.amount, 0) + COALESCE(sp.amount, 0)+COALESCE(cup.amount, 0) + COALESCE(sup.amount, 0))*0.1276,0) AS ST
FROM blocks AS b
JOIN clusters AS c
    ON b.block_code = c.block_code
JOIN schools AS s
    ON c.cluster_code = s.cluster_code
LEFT JOIN school_banks AS sb
    ON s.school_code = sb.school_code
LEFT JOIN cp
    ON s.school_code = cp.school_code
LEFT JOIN sp
   ON s.school_code = sp.school_code
LEFT JOIN cup
	ON s.school_code = cup.school_code
LEFT JOIN sup
	ON s.school_code = sup.school_code;
