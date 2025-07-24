package product

type product struct {
}

/*
เพื่อ ดึงข้อมูลสินค้า (product) และ คำตอบล่าสุด (ล่าสุดต่อคำถามของพนักงานแต่ละคน) จาก prn_product และ prn_answer โดยเฉพาะของพนักงานที่มีรหัส emp_code = '100222'

SELECT p.ProductID, a.*  									=> เลือกเอา ProductID จากตาราง prn_product (ตัวย่อคือ p) และเลือกทุกคอลัมน์จากตาราง prn_answer (ตัวย่อคือ a)
FROM prn_product p											=> เชื่อม (JOIN) ตาราง prn_product กับ prn_answer โดยใช้ ProductID ที่ตรงกัน (ใช้ INNER JOIN หมายถึงต้องมีข้อมูลในทั้งสองตารางถึงจะแสดงผล)
JOIN prn_answer a ON p.ProductID = a.ProductID				=> ^^^^^^^^^^
JOIN (														=>vvvvvvvvvvvvvvvvvvvvvvv
    SELECT emp_code, questionID, MAX(LogUp) AS MaxLogUp		=>	เลือก emp_code, questionID และเวลา LogUp ล่าสุด (ใช้ MAX(LogUp))
    FROM prn_answer											=>	จัดกลุ่มด้วย GROUP BY emp_code, questionID ⇒ ได้ "คำตอบล่าสุดของแต่ละคำถามต่อพนักงานแต่ละคน"
    GROUP BY emp_code, questionID							=>
) latest													=> ^^^^^^^^^^^^^^^^^^^^^^
ON a.emp_code = latest.emp_code								=> เชื่อม prn_answer กับผลลัพธ์ของ latest โดยระบุว่า: พนักงาน (emp_code) ต้องตรงกัน คำถาม (questionID) ต้องตรงกัน
   AND a.questionID = latest.questionID						=> และเวลาที่ตอบ (LogUp) ต้องตรงกับเวลาล่าสุดที่หาไว้ ทำให้เราดึงเฉพาะคำตอบที่ "ล่าสุดของแต่ละคำถาม" มาเท่านั้น
   AND a.LogUp = latest.MaxLogUp							=> ^^^^^^^^^^^^^^^^^^^^^^
WHERE p.emp_code = '100222';								=> กรองข้อมูลสินค้า (จาก prn_product) ให้แสดงเฉพาะสินค้าของพนักงาน emp_code = '100222'

คำสั่งนี้ดึงข้อมูลคำตอบที่ ใหม่ล่าสุดต่อคำถามของพนักงาน 100222 โดยจับคู่กับสินค้า (ProductID) ที่พนักงานเกี่ยวข้องด้วย
ถ้าคุณต้องการ:
เฉพาะคำตอบล่าสุด "โดยไม่สนคำถาม"
หรือเปลี่ยนเงื่อนไขไปยังพนักงานอื่น

*/

/*
SELECT p.ProductID , c.*
FROM prn_product p
JOIN prn_comm c on p.ProductID = c.ProductIDs
JOIN (
	SELECT emp_code , questionID , MAX(LogUp) as MaxLogUp
    FROM prn_comm
    GROUP BY emp_code , questionID
) lastest
on c.emp_code = lastest.emp_code
AND c.questionID = lastest.questionID
AND c.LogUp = lastest.MaxLogUp
WHERE p.emp_code='100222';

*/
